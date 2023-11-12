// Package domain definiert die Kern-Domänenmodelle und -Logik der Anwendung.
// In diesem Paket befindet sich die Definition des Beleg-Modells, das eine zentrale Rolle in der Geschäftslogik spielt.
//
// Der Beleg ist eine Domänen-Entität, die die wesentlichen Attribute und Geschäftsregeln
// für einen Beleg in der Anwendung repräsentiert. Diese Entität ist so gestaltet, dass sie
// die realen Geschäftsdaten und -prozesse so genau wie möglich widerspiegelt.
//
// Die Beleg-Entität enthält alle relevanten Informationen und Verhaltensweisen, die mit einem Beleg verbunden sind.
// Dies umfasst:
// - Attribute wie ID, Preis, Mehrwertsteuer (MwSt), Datum und Shop.
// - Methoden oder Funktionen, die Geschäftsoperationen oder Berechnungen darstellen,
//   die spezifisch für einen Beleg sind (falls erforderlich).
//
// Die Struktur und Logik in diesem Paket sollten unabhängig von technischen Aspekten wie Datenbank,
// Benutzeroberfläche oder externen Diensten sein. Dies unterstützt das Prinzip des Domain-Driven Designs,
// bei dem die Geschäftslogik und -regeln im Zentrum der Anwendungsentwicklung stehen.

package domain

type Beleg struct {
	Id    int     `json:"id,omitempty"`
	Price float64 `json:"price"`
	Mwst  float64 `json:"mwst"`
	Date  string  `json:"date"`
	Shop  string  `json:"shop"`
}
