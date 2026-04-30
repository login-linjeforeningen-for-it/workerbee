package internal

const (
	BASE_PATH                     = "/api/v2"
	USERINFO_URL                  = "https://authentik.login.no/application/o/userinfo/"
	ADMIN_GROUP                   = "QueenBee"
	BUCKET_NAME                   = "beehive"
	REGION                        = "ams3"
	IMG_PATH                      = "img/"
	ALBUM_PATH                    = "albums/"
	PROTECTED_REQUESTS_PER_MINUTE = 25
	MaxAlbumImageSize             = 1 << 20
	MaxDimension                  = 2400
	ImageRatio                    = 2.5
	MaxImageUploadSize            = 50 << 20
	WebPImageQuality              = 82
)
