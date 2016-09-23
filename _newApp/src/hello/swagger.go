// @APIVersion 1.0.0
// @basePath /{APP_NAME}/v1
package hello

// @Title list
// @Description get hello
// @Accept  json
// @Router /hello [get]
func list() {}

// @Title geteSku
// @Description get sku
// @Accept  json
// @Param   Limit     query    string     true        "limit"
// @Param   Offset     query    string     true        "offset"
// @Param   SESSION_ID     header    string     true        "ssn"
// @Param   TOKEN_ID     header    string     true        "token"
// @Router /hello/update [get]
func geteSku() {}

// @Title remove
// @Description delete sku
// @Accept  json
// @Param   Sku     path    string     true        "sku to be removed"
// @Param   SESSION_ID     header    string     true        "ssn"
// @Param   TOKEN_ID     header    string     true        "token"
// @Router /hello [delete]
func remove() {}

// @Title add
// @Description add sku
// @Accept  json
// @Param   BodyParam     body    AddParam     true        "body"
// @Param   SESSION_ID    header    string     true        "ssn"
// @Param   TOKEN_ID     header    string     true        "token"
// @Router /hello [post]
func add() {}
