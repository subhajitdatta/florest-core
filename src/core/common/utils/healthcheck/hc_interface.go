package healthcheck

type HCInterface interface {
	GetName() string
	GetHealth() map[string]interface{}
}
