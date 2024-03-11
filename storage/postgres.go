package storage

import (
	"clone/lms_back/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Store struct {
	DB      *sql.DB
	Branche brancheRepo
	Teacher teacherRepo
	Group groupRepo
	Student studentRepo
}

func New(cfg config.Config) (Store, error) {
	url := fmt.Sprintf(`host=%s port=%v user=%s password=%s database=%s sslmode=disable`,
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabase)

	db, err := sql.Open("postgres", url)
	if err != nil {
		return Store{}, err
	}

	newBranche := NewBranche(db)
	newTeacher := NewTeacher(db)
	newGroup:=NewGroup(db)
	newStudent:=NewStudent(db)

	return Store{
		DB:      db,
		Branche: newBranche,
		Teacher: newTeacher,
		Group:newGroup,
		Student:newStudent,
	}, nil

}
