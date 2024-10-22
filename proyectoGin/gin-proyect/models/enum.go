package models

type Status string

const(
	StatusCompletado   Status = "COMPLETED"
	StatusEnProceso    Status = "PENDING"
	StatusNoEscogida   Status = "FREE"
)

func StatusValido(status Status) bool{
	switch status{
	case StatusCompletado, StatusEnProceso, StatusNoEscogida:
		return true

	default:
		return false
	}	
}