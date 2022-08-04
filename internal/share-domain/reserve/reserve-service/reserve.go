package reserve_service

// import (
// 	user_reserve "example/restaurant-reserved/internal/infrastructure/repositories/user-reserve"
// 	"log"
// )

// type IReserveService interface {
// 	GetReserves() (string, error)
// }

// type reserveService struct {
// 	userReserveRepo user_reserve.IUserReserveRepository
// }

// func New(
// 	userReserveRepo user_reserve.IUserReserveRepository,
// ) IReserveService {
// 	return reserveService{
// 		userReserveRepo,
// 	}
// }

// // func (rs reserveService) GetReserves() (string, error) {
// // 	_, err := rs.userReserveRepo.GetReserves()
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}

// // 	return "model", nil
// // }
