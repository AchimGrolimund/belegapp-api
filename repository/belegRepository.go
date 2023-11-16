// Package repository beinhaltet die Implementierung des BelegRepository.
// Dieses Repository ist verantwortlich für alle Dateninteraktionen bezüglich der Beleg-Entitäten.
// Es abstrahiert die Details der Datenbankschicht von den Service- und Handler-Schichten.
//
// Das BelegRepository stellt die grundlegenden CRUD-Operationen (Create, Read, Update, Delete)
// für Beleg-Objekte bereit und kann für komplexere datenbankbezogene Operationen erweitert werden.
//
// Jegliche Interaktion mit der Beleg-Tabelle in der Datenbank sollte über dieses Repository erfolgen.
// Direkte Datenbankoperationen in anderen Teilen des Codes, wie Services oder Controllern, sollten vermieden werden,
// um die Einhaltung der DDD-Prinzipien (Domain-Driven Design) zu gewährleisten und die Wartbarkeit zu verbessern.
//
// Das BelegRepository sollte ausschließlich logische Operationen im Zusammenhang mit Beleg-Entitäten durchführen.
// Es sollte frei von Geschäftslogik sein, die stattdessen in der Service-Schicht implementiert wird.

package repository

import (
	"beleg-app/api/domain"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BelegRepositoryInterface interface {
	GetBelegById(ctx context.Context, id int) (*domain.Beleg, error)
	CreateBeleg(ctx context.Context, beleg *domain.Beleg) error
	DeleteBelegById(ctx context.Context, id int) error
	GetAllBelege(ctx context.Context) ([]domain.Beleg, error)
}

type BelegRepository struct {
	db *pgxpool.Pool
}

func NewBelegRepository(db *pgxpool.Pool) *BelegRepository {
	return &BelegRepository{db: db}
}

func (r *BelegRepository) GetBelegById(ctx context.Context, id int) (*domain.Beleg, error) {
	query := `SELECT "id", "b_price", "b_mwst", "b_date", "b_shop" FROM "belege" WHERE "id" = $1;`
	row := r.db.QueryRow(ctx, query, id)

	var beleg domain.Beleg
	if err := row.Scan(&beleg.Id, &beleg.Price, &beleg.Mwst, &beleg.Date, &beleg.Shop); err != nil {
		return nil, err
	}

	return &beleg, nil
}

func (r *BelegRepository) CreateBeleg(ctx context.Context, beleg *domain.Beleg) error {
	query := `INSERT INTO "belege"("b_price", "b_mwst", "b_date", "b_shop") VALUES ($1, $2, $3, $4);`
	_, err := r.db.Exec(ctx, query, beleg.Price, beleg.Mwst, beleg.Date, beleg.Shop)
	return err
}

func (r *BelegRepository) DeleteBelegById(ctx context.Context, id int) error {
	query := `DELETE FROM "belege" WHERE "id" = $1;`
	_, err := r.db.Exec(ctx, query, id)
	return err
}

func (r *BelegRepository) GetAllBelege(ctx context.Context) ([]domain.Beleg, error) {
	query := `SELECT "id", "b_price", "b_mwst", "b_date", "b_shop" FROM "belege";`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var belege []domain.Beleg
	for rows.Next() {
		var beleg domain.Beleg
		if err := rows.Scan(&beleg.Id, &beleg.Price, &beleg.Mwst, &beleg.Date, &beleg.Shop); err != nil {
			return nil, err
		}
		belege = append(belege, beleg)
	}

	return belege, nil
}
