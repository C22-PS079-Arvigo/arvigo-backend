package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/yusufwib/arvigo-backend/constant"
	"github.com/yusufwib/arvigo-backend/datastruct"
	"github.com/yusufwib/arvigo-backend/utils"
	"gorm.io/gorm"
)

func CreateInitialProduct(data datastruct.CreateInitialProductInput) (statusCode int, err error) {
	statusCode = http.StatusCreated

	var (
		db                 = Database()
		currentTime        = time.Now()
		imagesURL          []string
		detailVariants     []datastruct.DetailProductVariant
		productTagsPayload []datastruct.DetailProductTag
	)

	err = json.Unmarshal([]byte(data.DetailProductVariants), &detailVariants)
	if err != nil {
		return http.StatusBadRequest, errors.New("failed to parse variants")
	}

	for _, img := range data.Images {
		url, err := UploadImageToGCS(img)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		imagesURL = append(imagesURL, url)
	}

	productPayload := datastruct.Product{
		Name:         data.Name,
		Description:  data.Description,
		Images:       strings.Join(imagesURL, ","),
		LinkExternal: "", // coming soon!
		CategoryID:   data.CategoryID,
		BrandID:      data.BrandID,
		MerchantID:   0, // 0 is for create from admin
		CreatedAt:    currentTime,
		UpdatedAt:    currentTime,
	}

	// Begin a transaction
	tx := db.Begin()

	// Defer the rollback function in case of error
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = tx.Create(&productPayload).Error; err != nil {
		return http.StatusInternalServerError, err
	}

	explodedTags := strings.Split(data.DetailProductTags, ",")
	for _, tagID := range explodedTags {
		tagIDNum := utils.StrToUint64(tagID, 0)
		if tagIDNum == 0 {
			return http.StatusInternalServerError, errors.New("cannot add tags")
		}

		productTagsPayload = append(productTagsPayload, datastruct.DetailProductTag{
			TagID:     tagIDNum,
			ProductID: productPayload.ID,
			CreatedAt: currentTime,
			UpdatedAt: currentTime,
		})
	}

	if err = tx.Create(&productTagsPayload).Error; err != nil {
		return http.StatusInternalServerError, err
	}

	// Create variants
	for i := range detailVariants {
		detailVariants[i].ProductID = productPayload.ID
		detailVariants[i].CreatedAt = currentTime
		detailVariants[i].UpdatedAt = currentTime
	}
	if err = tx.Create(&detailVariants).Error; err != nil {
		return http.StatusInternalServerError, err
	}

	// Commit the transaction if all queries succeed
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		log.Println("Error committing transaction:", err)
		return http.StatusInternalServerError, err
	}

	return
}

func CreateMerchantProduct(data datastruct.CreateMerchantProductInput) (statusCode int, err error) {
	statusCode = http.StatusCreated

	var (
		db                    = Database()
		currentTime           = time.Now()
		imagesURL             []string
		initialProduct        datastruct.Product
		linkedProduct         datastruct.DetailLinkedProduct
		initialProductVariant []datastruct.DetailProductVariant
		initialProductTag     []datastruct.DetailProductTag
		detailMarketplaces    []datastruct.DetailProductMarketplace
	)

	err = json.Unmarshal([]byte(data.DetailProductMarketplace), &detailMarketplaces)
	if err != nil {
		return http.StatusBadRequest, errors.New("failed to parse marketplaces")
	}

	for _, img := range data.Images {
		url, err := UploadImageToGCS(img)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		imagesURL = append(imagesURL, url)
	}

	// select initial products
	if err = db.Table("products").
		Select("*").
		Find(&initialProduct).
		Where("id = ?", data.ProductID).
		Error; err != nil {
		return http.StatusInternalServerError, err
	}

	if initialProduct.ID == 0 {
		return http.StatusNotFound, errors.New("initial products not found")
	}

	// Begin a transaction
	tx := db.Begin()

	// Defer the rollback function in case of error
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// remap payload
	initialProduct = datastruct.Product{
		ID:           0,
		Name:         data.Name,
		Description:  data.Description,
		Images:       strings.Join(imagesURL, ","),
		LinkExternal: "", // coming soon
		CategoryID:   initialProduct.CategoryID,
		BrandID:      initialProduct.BrandID,
		MerchantID:   data.MerchantID,
		Price:        data.Price,
		CreatedAt:    currentTime,
		UpdatedAt:    currentTime,
	}

	if err = tx.Create(&initialProduct).Error; err != nil {
		return http.StatusInternalServerError, err
	}

	// select initial products variants
	if err = db.Table("detail_product_variants").
		Select("*").
		Where("product_id = ?", data.ProductID).
		Find(&initialProductVariant).
		Error; err != nil {
		return http.StatusInternalServerError, err
	}

	if len(initialProductVariant) == 0 {
		return http.StatusNotFound, errors.New("initial product variants not found")
	}

	for i, variant := range initialProductVariant {
		initialProductVariant[i] = datastruct.DetailProductVariant{
			ID:               0,
			Name:             variant.Name,
			LinkAR:           variant.LinkAR,
			IsPrimaryVariant: variant.IsPrimaryVariant,
			ProductID:        initialProduct.ID,
			CreatedAt:        currentTime,
			UpdatedAt:        currentTime,
		}
	}

	if err = tx.Create(&initialProductVariant).Error; err != nil {
		return http.StatusInternalServerError, err
	}

	// select initial product tags
	if err = db.Table("detail_product_tags").
		Select("*").
		Where("product_id = ?", data.ProductID).
		Find(&initialProductTag).
		Error; err != nil {
		return http.StatusInternalServerError, err
	}

	if len(initialProductTag) == 0 {
		return http.StatusNotFound, errors.New("initial product tags not found")
	}

	for i, tag := range initialProductTag {
		initialProductTag[i] = datastruct.DetailProductTag{
			ID:        0,
			TagID:     tag.TagID,
			ProductID: initialProduct.ID,
			CreatedAt: currentTime,
			UpdatedAt: currentTime,
		}
	}

	if err = tx.Create(&initialProductTag).Error; err != nil {
		return http.StatusInternalServerError, err
	}

	for i, marketplace := range detailMarketplaces {
		detailMarketplaces[i] = datastruct.DetailProductMarketplace{
			MarketplaceID:   marketplace.ID,
			ParentProductID: data.ProductID,
			ProductID:       initialProduct.ID,
			AddressID:       marketplace.AddressID,
			Link:            marketplace.Link,
			Clicked:         0,
			CreatedAt:       currentTime,
			UpdatedAt:       currentTime,
		}
	}

	if err = tx.Create(&detailMarketplaces).Error; err != nil {
		return http.StatusInternalServerError, err
	}

	linkedProduct = datastruct.DetailLinkedProduct{
		InitialProductID:  data.ProductID,
		MerchantProductID: initialProduct.ID,
		MerchantID:        data.MerchantID,
		CreatedAt:         time.Time{},
		UpdatedAt:         time.Time{},
	}

	if err = tx.Create(&linkedProduct).Error; err != nil {
		return http.StatusInternalServerError, err
	}

	// Commit the transaction if all queries succeed
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		log.Println("Error committing transaction:", err)
		return http.StatusInternalServerError, err
	}

	return
}

func GetInitialProductByCategoryID(categoryID uint64) (res []datastruct.InitialProductResponse, statusCode int, err error) {
	db := Database()
	statusCode = http.StatusOK

	var (
		productIDs        []uint64
		products          []datastruct.InitialProduct
		productVariants   []datastruct.InitialProductVariant
		productVariantMap = make(map[uint64][]datastruct.InitialProductVariant, 0)
	)
	if err := db.Table("products p").
		Select([]string{
			"p.id",
			"p.name",
			"p.description",
			"p.images",
			"p.link_external",
			"c.name as category_name",
			"b.name as brand_name",
		}).
		Where("p.merchant_id = 0 AND p.category_id = ?", categoryID).
		Joins("LEFT JOIN categories c on c.id = p.category_id").
		Joins("LEFT JOIN brands b on b.id = p.brand_id").
		Find(&products).
		Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	if len(products) == 0 {
		return res, http.StatusNotFound, errors.New("products not found")
	}

	//collect product ids
	for _, product := range products {
		productIDs = append(productIDs, product.ID)
	}

	if err := db.Table("detail_product_variants").
		Select([]string{
			"name",
			"link_ar",
			"is_primary_variant",
			"product_id",
		}).
		Where("product_id IN (?) ", productIDs).
		Find(&productVariants).
		Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	for _, variant := range productVariants {
		productVariantMap[variant.ProductID] = append(productVariantMap[variant.ProductID], variant)
	}

	for _, product := range products {
		res = append(res, datastruct.InitialProductResponse{
			InitialProduct: product,
			Images:         strings.Split(product.Images, ","),
			Variants:       productVariantMap[product.ID],
		})
	}

	return
}

func GetInitialProductByID(productID uint64) (res datastruct.InitialProductResponse, statusCode int, err error) {
	db := Database()
	statusCode = http.StatusOK

	var (
		marketplaceDetailIDs []uint64
		products             datastruct.InitialProduct
		productVariants      []datastruct.InitialProductVariant
	)
	if err := db.Table("products p").
		Select([]string{
			"p.id",
			"p.name",
			"p.description",
			"p.images",
			"p.link_external",
			"c.name as category_name",
			"b.name as brand_name",
		}).
		Where("p.merchant_id = 0 AND p.id = ?", productID).
		Joins("LEFT JOIN categories c on c.id = p.category_id").
		Joins("LEFT JOIN brands b on b.id = p.brand_id").
		First(&products).
		Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	if err := db.Table("detail_product_variants").
		Select([]string{
			"name",
			"link_ar",
			"is_primary_variant",
			"product_id",
		}).
		Where("product_id = ? ", productID).
		Find(&productVariants).
		Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	if err := db.Table("detail_linked_products").
		Select([]string{
			"dpm.id",
		}).
		Joins("left join detail_product_marketplaces dpm on detail_linked_products.merchant_product_id = dpm.product_id").
		Where("initial_product_id = ? ", productID).
		Find(&marketplaceDetailIDs).
		Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	var merchantProduct []datastruct.ProductMarketplaceWishlist
	if len(marketplaceDetailIDs) > 0 {

		if err = db.Table("detail_product_marketplaces").
			Select([]string{
				"detail_product_marketplaces.id AS id",
				"products.name",
				"brands.name AS brand",
				"products.images",
				"products.price",
				"merchants.name AS merchant",
				"detail_product_marketplaces.link AS marketplace_link",
				"detail_product_marketplaces.marketplace_id",
				"detail_product_marketplaces.addresses_id",
			}).
			Joins("LEFT JOIN products ON products.id = detail_product_marketplaces.product_id").
			Joins("LEFT JOIN brands ON brands.id = products.brand_id").
			Joins("LEFT JOIN merchants ON products.merchant_id = merchants.id").
			Where("detail_product_marketplaces.id IN (?)", marketplaceDetailIDs).
			Find(&merchantProduct).Error; err != nil {
			return res, http.StatusInternalServerError, err
		}

		for i, v := range merchantProduct {
			merchantProduct[i].Image = strings.Split(v.Image, ",")[0]

			if v.AddressID != 0 {
				merchantProduct[i].Type = "offline"
				addr, _, _, err := GetAddressByID(v.AddressID)
				if err == nil {
					merchantProduct[i].Address = &addr
				}
				continue
			}

			if v.MarketplaceID != 0 {
				merchantProduct[i].Type = "online"
				marketplaceName := constant.Marketplace[v.MarketplaceID]
				merchantProduct[i].Marketplace = &marketplaceName
			}
		}
	}

	var tagIDs []uint64
	var tags []string
	if err := db.Table("detail_product_tags").
		Select([]string{
			"tag_id",
		}).
		Where("product_id = ?", productID).
		Find(&tagIDs).
		Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	for _, v := range tagIDs {
		tags = append(tags, constant.GetTagNameByDetailTag[v]...)
	}

	res = datastruct.InitialProductResponse{
		InitialProduct:  products,
		Images:          strings.Split(products.Images, ","),
		Variants:        productVariants,
		ListMarketplace: merchantProduct,
		Tags:            utils.RemoveDuplicates(tags),
	}

	return
}

func GetMarketplaceProductByID(productID uint64) (merchantProduct datastruct.ProductMarketplaceDetail, statusCode int, err error) {
	db := Database()
	statusCode = http.StatusOK

	if err = db.Table("detail_product_marketplaces").
		Select([]string{
			"detail_product_marketplaces.id AS id",
			"products.name",
			"brands.name AS brand",
			"products.images",
			"products.price",
			"merchants.name AS merchant",
			"detail_product_marketplaces.link AS marketplace_link",
			"detail_product_marketplaces.marketplace_id",
			"detail_product_marketplaces.addresses_id",
		}).
		Joins("LEFT JOIN products ON products.id = detail_product_marketplaces.product_id").
		Joins("LEFT JOIN brands ON brands.id = products.brand_id").
		Joins("LEFT JOIN merchants ON products.merchant_id = merchants.id").
		Where("detail_product_marketplaces.id = ?", productID).
		Find(&merchantProduct).Error; err != nil {
		return merchantProduct, http.StatusInternalServerError, err
	}

	merchantProduct.Images = strings.Split(merchantProduct.Image, ",")
	if merchantProduct.AddressID != 0 {
		merchantProduct.Type = "offline"
		addr, _, _, err := GetAddressByID(merchantProduct.AddressID)
		if err == nil {
			merchantProduct.Address = &addr
		}
	} else if merchantProduct.MarketplaceID != 0 {
		merchantProduct.Type = "online"
		marketplaceName := constant.Marketplace[merchantProduct.MarketplaceID]
		merchantProduct.Marketplace = &marketplaceName
	}

	if err = db.Table("detail_product_marketplaces").
		Where("id = ?", productID).
		UpdateColumn("clicked", gorm.Expr("clicked + ?", 1)).Error; err != nil {
		return merchantProduct, http.StatusInternalServerError, err
	}

	return
}

func GetProductRecommendationMachineLearning() (res []datastruct.ProductRecommendationResponse, statusCode int, err error) {
	statusCode = http.StatusOK

	var (
		db               = Database()
		linkedProductIDs = []uint64{}
		merchantIDs      = []uint64{}
	)

	columns := []string{
		"p.id",
		"p.name",
		"p.description",
		"c.name AS category",
		"b.name AS brand",
		"GROUP_CONCAT(DISTINCT t.name) AS tags",
		"GROUP_CONCAT(DISTINCT dlp.merchant_id) AS merchant_id",
		"GROUP_CONCAT(DISTINCT dlp.merchant_product_id) AS linked_product",
	}

	if err = db.Table("products p").Select(columns).
		Joins("LEFT JOIN categories c ON p.category_id = c.id").
		Joins("LEFT JOIN brands b ON p.brand_id = b.id").
		Joins("LEFT JOIN detail_product_tags dpt ON p.id = dpt.product_id").
		Joins("LEFT JOIN tags t ON dpt.tag_id = t.id").
		Joins("LEFT JOIN detail_linked_products dlp ON p.id = dlp.initial_product_id").
		Where("p.merchant_id = ?", 0).
		Group("p.id").
		Find(&res).Error; err != nil {
		return res, http.StatusInternalServerError, err
	}

	for _, v := range res {
		merchantIDSplitted := strings.Split(v.MerchantIDs, ",")
		linkedIDSplitted := strings.Split(v.ProductIDs, ",")

		for _, id := range merchantIDSplitted {
			merchantID := utils.StrToUint64(id, 0)
			merchantIDs = append(merchantIDs, merchantID)
		}

		for _, id := range linkedIDSplitted {
			productID := utils.StrToUint64(id, 0)
			linkedProductIDs = append(linkedProductIDs, productID)
		}
	}

	return
}

func GetProductRecommendationMachineLearningDummy() (res []datastruct.ProductRecommendationResponse, statusCode int, err error) {
	tags := []string{"circle", "heart", "oblong", "oval", "square", "triangle"}
	brands := []string{"RayBan", "Oakley", "Baleno", "CHANEL", "Police", "Emporio"}
	merchants := []string{"Optik Susi", "Optik Sukarno", "Optik Merah Putih", "Optik tik"}

	// Generate 10 dummy data entries
	for i := 1; i <= 20; i++ {
		product := datastruct.ProductRecommendationResponse{
			ID:          uint64(i),
			Name:        fmt.Sprintf("Kacamata %d", i),
			Description: fmt.Sprintf("This is the description of Product %d", i),
			Category:    "Glasses",
			Brand:       strings.Join(getRandomTags(brands, 1), ", "),
			Tags:        strings.Join(getRandomTags(tags, 3), ", "),
			Merchants:   strings.Join(getRandomTags(merchants, 4), ", "),
			Clicked:     uint64(generateRandomTags()),
		}
		res = append(res, product)
	}

	return
}

func getRandomTags(tags []string, count int) []string {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(tags), func(i, j int) {
		tags[i], tags[j] = tags[j], tags[i]
	})
	return tags[:count]
}

func generateRandomTags() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 100
	return rand.Intn(max-min+1) + min
}

func VerifyMerchantProduct(data datastruct.VerifyProductInput) (statusCode int, err error) {
	db := Database()
	statusCode = http.StatusOK

	if err = db.Model(&datastruct.Product{}).
		Where("id = ?", data.ProductID).
		Updates(map[string]interface{}{
			"status":        data.Status,
			"rejected_note": data.RejectedNote,
			"updated_at":    time.Now(),
		}).Error; err != nil {
		return http.StatusInternalServerError, err
	}

	return
}

func UpdateMerchantProduct(data datastruct.UpdateProductInput) (statusCode int, err error) {
	db := Database()
	statusCode = http.StatusOK

	if err = db.Model(&datastruct.Product{}).
		Where("id = ?", data.ProductID).
		Updates(map[string]interface{}{
			"price":       data.Price,
			"description": data.Description,
			"updated_at":  time.Now(),
		}).Error; err != nil {
		return http.StatusInternalServerError, err
	}

	return
}
