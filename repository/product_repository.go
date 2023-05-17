package repository

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/yusufwib/arvigo-backend/datastruct"
	"github.com/yusufwib/arvigo-backend/utils"
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
		Name:         initialProduct.Name,
		Description:  data.Description,
		Images:       strings.Join(imagesURL, ","),
		LinkExternal: "", // coming soon
		CategoryID:   initialProduct.CategoryID,
		BrandID:      initialProduct.BrandID,
		MerchantID:   data.MerchantID,
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
			MarketplaceID: marketplace.ID,
			ProductID:     initialProduct.ID,
			AddressID:     marketplace.AddressID,
			Link:          marketplace.Link,
			Clicked:       0,
			CreatedAt:     currentTime,
			UpdatedAt:     currentTime,
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