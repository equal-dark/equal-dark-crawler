package crawlers

import "errors"

// ErrorInvalidBrandID is not registered brand id
var ErrorInvalidBrandID = errors.New("invalid brandID")

// ErrorInvalidPage is catched error page
var ErrorInvalidPage = errors.New("invalid products page")
