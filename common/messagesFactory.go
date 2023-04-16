package common

type Resource string

const (
	UserResource       Resource = "User"
	RestaurantResource          = "Restaurante"
)

type SuccesMessage struct {
	Message Resource
}

func CreateSuccesCreatedMessage(resource Resource) SuccesMessage {
	return SuccesMessage{Message: resource + " created successfully"}
}
