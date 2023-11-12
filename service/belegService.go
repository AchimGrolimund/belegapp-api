// Package service beinhaltet die Implementierung des BelegService.
// Der BelegService ist zentraler Bestandteil der Geschäftslogikschicht und verantwortlich für die Anwendungslogik,
// die sich auf die Beleg-Entitäten bezieht. Dieser Service agiert als Mittler zwischen den HTTP-Handlern und dem Daten-Repository.
//
// In BelegService werden Geschäftsregeln und Logik implementiert, die mit Beleg-Operationen assoziiert sind.
// Diese Schicht stellt sicher, dass die Daten, die vom Repository kommen, entsprechend den Geschäftsregeln verarbeitet und
// manipuliert werden, bevor sie an die Handler-Schicht weitergeleitet oder von dort abgerufen werden.
//
// Die Hauptaufgaben des BelegService umfassen:
// - Validierung von Eingabedaten
// - Ausführung von Geschäftsregeln
// - Koordination von Transaktionen
// - Zusammenarbeit mit anderen Services, falls erforderlich, um komplexe Geschäftsvorgänge durchzuführen
//
// Es ist wichtig, dass der BelegService von direkten Datenbankoperationen abstrahiert.
// Alle Datenbankinteraktionen sollten durch das BelegRepository erfolgen, um die Trennung der Verantwortlichkeiten
// gemäß den DDD-Prinzipien (Domain-Driven Design) zu gewährleisten.

package service

import (
	"beleg-app/api/domain"
	"beleg-app/api/repository"
	"context"
)

type BelegService struct {
	repo *repository.BelegRepository
}

func NewBelegService(repo *repository.BelegRepository) *BelegService {
	return &BelegService{repo: repo}
}

func (s *BelegService) GetBelegById(ctx context.Context, id int) (*domain.Beleg, error) {
	return s.repo.GetBelegById(ctx, id)
}

func (s *BelegService) CreateBeleg(ctx context.Context, beleg *domain.Beleg) error {
	// Hier können Sie zusätzliche Geschäftslogik einfügen, z.B. Validierung der Belegdaten

	return s.repo.CreateBeleg(ctx, beleg)
}

func (s *BelegService) DeleteBelegById(ctx context.Context, id int) error {
	return s.repo.DeleteBelegById(ctx, id)
}

func (s *BelegService) GetAllBelege(ctx context.Context) ([]domain.Beleg, error) {
	return s.repo.GetAllBelege(ctx)
}
