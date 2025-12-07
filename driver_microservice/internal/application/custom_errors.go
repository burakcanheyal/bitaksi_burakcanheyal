package application

var CustomErrorMap = map[string]struct {
	Status  int
	Message string
}{
	// Generic
	"ERR_BAD_REQUEST": {400, "Invalid request"},
	"ERR_INTERNAL":    {500, "Internal server error"},
	"Kayıtlı Plaka":   {400, "Plaka zaten kayıtlı"},

	// Driver Specific
	"ERR_MISSING_ID":         {400, "Id girmek zorunlu!"},
	"Hatalı ID":              {400, "Hatalı Id"},
	"Sürücü Bulunamadı":      {404, "Sürücü Bulunamadı"},
	"ERR_VALIDATION":         {400, "Validation error"},
	"Koordinat Hatası : Lon": {400, "Lon koordinatı validasyon hatası"},
	"Koordinat Hatası: Lat":  {400, "Lat koordinatı validasyon hatası"},
	"Taksi Tipi Hatası":      {400, "Taksi Tipi hatası"},
	"lastName":               {400, "Soyisim en az 2 karakter olmalı"},
	"firstName":              {400, "İsim en az 2 karakter olmalı"},
	"pageSize":               {400, "pageSize 1 ile 100 arasında olmalı"},
	"page":                   {400, "Page 1'den büyük olmalı"},
}

type ApiError struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}

func Wrap(code string) *ApiError {
	if val, ok := CustomErrorMap[code]; ok {
		apiErr := ApiError{Code: val.Status, Msg: val.Message}
		return &apiErr
	}
	return nil
}

func (e *ApiError) Error() string {
	return string(e.Code)
}
