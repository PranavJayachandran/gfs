package masterInfrastructure

import (
	"fmt"

	"github.com/google/uuid"
)

func getUniqueFileName() string {
	uniqueID := uuid.New().String()
	fileName := fmt.Sprintf("file_%s.txt", uniqueID)
	return fileName
}
