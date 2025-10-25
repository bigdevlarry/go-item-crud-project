package helpers

import (
	"go-test/backend/models/dto"
	"go-test/backend/models/enums"
	"strings"
)

func NormalizeCreateDTO(createDTO *dto.ItemCreateDTO) {
	createDTO.Type = enums.ItemType(strings.ToUpper(string(createDTO.Type)))
	createDTO.Status = enums.ItemStatus(strings.ToUpper(string(createDTO.Status)))
}

func NormalizeUpdateDTO(updateDTO *dto.ItemUpdateDTO) {
	if updateDTO.Type != nil {
		*updateDTO.Type = enums.ItemType(strings.ToUpper(string(*updateDTO.Type)))
	}
	if updateDTO.Status != nil {
		*updateDTO.Status = enums.ItemStatus(strings.ToUpper(string(*updateDTO.Status)))
	}
}
