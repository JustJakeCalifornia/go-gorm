package repository

// 1. Variante
// repository pattern
// fuer jede domain wie z.b. user, animal etc.
// wuerde man ein Repository wie hier machen.
// type UserRepository struct {
// 	db *sql.DB
// }

// func NewUserRepository() *UserRepository {
// 	return &UserRepository{
// 		db: GetDatabase(),
// 	}
// }

// 2. Variante
// Singleton pattern
// var lock = &sync.Mutex{}

// type Manager struct {
// 	DB *gorm.DB
// }

// var singleInstance *Manager

// func GetInstance() *Manager {
// 	if singleInstance == nil {
// 		lock.Lock()
// 		defer lock.Unlock()
// 		if singleInstance == nil {
// 			fmt.Println("Creting Single Instance Now")
// 			singleInstance = &Manager{DB: getDatabase()}
// 		} else {
// 			fmt.Println("Single Instance already created-1")
// 		}
// 	} else {
// 		fmt.Println("Single Instance already created-2")
// 	}
// 	return singleInstance
// }
