# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).


## [Released]
### [1.0.0] - yusufw2429@gmail.com
#### Added
- Add JWT for authentication
- Add seeders for indonesian province, district, sub-disctrict and postal code
- Add DDL and seeders
- Add login, register, get users and get user by id
- Add validator pkg
- Add response helper
- Add location module
- Add face shape module
- Add gcs bucket integration
- Add utils: random, base64, converter, response, validator, string, array

### [1.0.1] - yusufw2429@gmail.com
#### Added
- Update login response
- Update engine sql
- Add more validation

### [1.0.2] - yusufw2429@gmail.com
#### Added
- Graceful shutdown
- Add health check
- Prevent sql injection
- Create product initial
- Create product merchant
- Get product initial
- Upload img to gcs helper (reusable)
- Add categories module
- Add brands module
- Add questionnaires module
- Add verify product from admin
- Update cred.json

### [1.0.3] - yusufw2429@gmail.com
#### Fixed
- Fix get partners and users
- Add validation for register user and partner

### [1.0.4] - yusufw2429@gmail.com
#### Added
- Add wishlist
- Add parse token jwt to context
- Add price on product
- Tidy up routes
- Add update product on partner app
- Add dummy data for machine learning

### [1.0.5] - yusufw2429@gmail.com
#### Added
- Update face detection
- Add get wishlist

### [1.0.6] - yusufw2429@gmail.com
#### Added
- Add home page
- Add constant value for homepage
- Update DDL

### [1.0.7] - yusufw2429@gmail.com
#### Added
- Add product search
- Add get product by brand
- Add get product by category
- Add home merchant


### [1.0.8] - yusufw2429@gmail.com
#### Added
- Add X-API-Key for integration with ML (not using jwt)
- Fixing bug
- Update deployment
- Add resty pkg utils
- Integrate questionnaire with ML
- Update user profile

## [Unreleased]
### [1.0.9] - yusufw2429@gmail.com
#### Added
- Delete wishlist
- Add home partner app
- Delete product
- Product by id on mitra
- Merchant and products
- subscription module