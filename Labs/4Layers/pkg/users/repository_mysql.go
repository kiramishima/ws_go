package users

import (
	"MenuDigital/pkg/entities"
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type MySQLRepository struct {
	conn *sqlx.DB
}

//NewMySQLRepository create new repository
func NewMySQLRepository(p *sqlx.DB) *MySQLRepository {
	return &MySQLRepository{
		conn: p,
	}
}

func (m *MySQLRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []entities.User, err error) {
	rows, err := m.conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]entities.User, 0)
	for rows.Next() {
		t := entities.User{}
		err = rows.Scan(
			&t.ID,
			&t.Email,
			&t.RoleID,
			&t.IsBlocked,
			&t.UpdatedAt,
			&t.CreatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

//Find a bookmark
func (r *MySQLRepository) Find(ctx context.Context, id int16) (*entities.User, error) {
	/*query := `SELECT * FROM users WHERE ID = ?`
	stmt, err := r.conn.PreparexContext(ctx, query)
	if err != nil {
		return nil, err
	}
	switch err {
	case nil:
		return &result, nil
	case mysql.ErrNotFound:
		return nil, entity.ErrNotFound
	default:
		return nil, err
	}*/
	return nil, nil
}

//Store a bookmark
func (r *MySQLRepository) Store(ctx context.Context, b *entities.User) (err error) {
	query := ""
	stmt, err := r.conn.PreparexContext(ctx, query)
	if err != nil {
		return
	}
	res, err := stmt.ExecContext(ctx, b.Email, b.Password)
	if err != nil {
		return
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	b.ID = lastID
	return
}

//FindAll users
func (r *MySQLRepository) FindAll(ctx context.Context) (res []entities.User, err error) {
	query := `SELECT * FROM users`
	rows, err := r.conn.QueryContext(ctx, query)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	res = make([]entities.User, 0)
	for rows.Next() {
		t := entities.User{}
		err = rows.Scan(
			&t.ID,
			&t.Email,
			&t.Password,
			&t.Token,
			&t.RoleID,
			&t.IsBlocked,
			&t.UpdatedAt,
			&t.CreatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		res = append(res, t)
	}

	return res, nil
}

//Search bookmarks
func (r *MySQLRepository) Search(ctx context.Context, q string) (result []entities.User, err error) {
	query := `SELECT * FROM users WHERE name = ?`
	rows, err := r.conn.QueryContext(ctx, query, q)

	if err != nil {
		//logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]entities.User, 0)
	for rows.Next() {
		t := entities.User{}
		err = rows.Scan(
			&t.ID,
			&t.Email,
			&t.RoleID,
			&t.IsBlocked,
			&t.UpdatedAt,
			&t.CreatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

//Delete an user
func (r *MySQLRepository) Delete(ctx context.Context, id int64) (err error) {
	query := "DELETE FROM users WHERE id = ?"

	stmt, err := r.conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return
	}

	if rowsAfected != 1 {
		err = fmt.Errorf("Weird  Behavior. Total Affected: %d", rowsAfected)
		return
	}

	return
}
