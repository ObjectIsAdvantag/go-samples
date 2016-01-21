// Copyright 2016, Stève Sfartz
// Licensed under the MIT License

// TRANSLATE  Ensemble de fonctions utilitaires pour manipuler des dates et âges
//
// Le type AnneesMoisJours permet d'exprimer des dates et âges en Années, Mois et Jour,
// Le type AnneesMoisJours peut être converti en type time.Time (UTC) et vice versa.
// Enfin, il est possible d'obtenir une durée en AnneesMoisJours à partir de la différence entre 2 dates.
package main

import (
	"time"
	"errors"
)

var ErrOverLimits error = errors.New("Over limits")
var ErrBadArguments error = errors.New("Bad function arguments")



// TRANSLATE  Le type AnneesMois représente une date sous la forme d'un nombre d'années, de mois et de jour
// Exemple : 62 ans et 7 mois
type YearsMonthsDays struct {
	Years  int `json:"annees,required"` // nombre d'années
	Months int `json:"mois,omitempty"`  // nombre de mois
	Days   int `json:"jours,omitempty"` // nombre de jours
}

// TRANSLATE  Retourne la valeur en années
func (ymd YearsMonthsDays) ToYears() float32 {
	return float32(ymd.Years) + float32(ymd.Months)/12 + float32(ymd.Days /365)
}

// TRANSLATE  Retourne la valeur en mois
func (ymd YearsMonthsDays) ToMonths() float32 {
	return float32(ymd.Years *12) + float32(ymd.Months) + float32(ymd.Days /365*12)
}

// TRANSLATE  Crée un nouvel objet de type time.Time pour l' AnneesMoisJours spécifié
func YearsMonthsDaysToTime(ymd YearsMonthsDays) (time.Time, error) {
	if ymd.Years < 0 || ymd.Years > 2100 {
		return time.Time{}, ErrOverLimits
	}
	if ymd.Months < 1 || ymd.Months > 12 {
		return time.Time{}, ErrOverLimits
	}
	if ymd.Days < 1 || ymd.Days > 32 {
		return time.Time{}, ErrOverLimits
	}

	return time.Date(ymd.Years, time.Month(ymd.Months), ymd.Days, 0, 0, 0, 0, time.UTC), nil
}

// TRANSLATE Crée un nouvel objet de type AnneesMoisJours pour la date spécifiée
func TimeToYearsMonthsDays(t time.Time) (YearsMonthsDays, error) {
	if t.IsZero() {
		return YearsMonthsDays{}, ErrBadArguments
	}

	return YearsMonthsDays{t.Year(), int(t.Month()), t.Day()}, nil
}


// This SubDate implementation leverages the AddDate function from the standard golang library
func SubDateStandard(date time.Time, years int, months int, days int) (time.Time, error) {
	return date.AddDate(years * -1, months * -1, days * -1), nil
}

func SubDateNew(date time.Time, years int, months int, days int) (time.Time, error) {

	// UPDATE 2016/1/21 La fonction Time.AddDate accepte des arguments négatifs !
	// Etudier s'il est possible de remplacer cet algo maison par :
	// comp, _ := jusque.AddDate(depuis.Year() * -1, int(depuis.Month()) * -1, depuis.Day() * -1)
	// return TimeToAnneesMoisJour(comp), nil

	// Formule de calcul de l'age en année / mois
	// soient AAAA2/MM2/DD2 - AAAA1/MM1/DD1
	//
	// 1. Se rendre à l'année cible et voir si on a dépassé
	// - si ce n'est pas le cas, ok
	// - si c'est le cas, revenir 1 an en arrière,
	// - mémoriser l'année cible, le bond réalisé en année, et le fait qu'on a dû ou non s'arrêter un an avant
	//
	// 2. Se rendre sur le mois cible candidat et calculer la différence de jours
	// - si elle est positive, ne rien faire, on peut réaliser l'opération
	// - si elle est négative, il faut opérer un changement de mois (avec le cas particulier du mois de janvier qui se transforme en décembre)
	// - mémoriser le mois cible et calculer le bond réalisé en mois
	//
	// 3. Se rendre sur l'année et mois cible, et calculer la différence en seconde entre les 2 dates
	// - convertir cette différence en jours

	// 1.
	tempDateCible, _ := YearsMonthsDaysToTime(YearsMonthsDays{
		Years: date.Year(),
		Months:   months,
		Days:  days,
	})

	anneeCible := date.Year()
	changementAnnee := 0
	if tempDateCible.After(date) {
		anneeCible--
		changementAnnee = 1
	}

	// 2.
	moisCible := int(date.Month())
	changementMois2 := 0
	changementAnnee2 := 0
	if days > date.Day() {
		moisCible--
		changementMois2 = 1
		if moisCible == 0 {
			moisCible = 12
			changementAnnee2 = 1
		}
	}
	nbMois := int(date.Month()) + 12*changementAnnee - months - changementMois2

	// 3.
	tempDateCible, _ = YearsMonthsDaysToTime(YearsMonthsDays{
		Years: date.Year() - changementAnnee2,
		Months:   moisCible,
		Days:  days,
	})
	deltaJours := date.Sub(tempDateCible).Minutes() / 60 / 24

	return time.Date(anneeCible - years, time.Month(nbMois), int(deltaJours), date.Hour(), date.Minute(), date.Second(), date.Nanosecond(), date.Location()), nil
}

func SubDate(substracted time.Time, date time.Time) (YearsMonthsDays, error) {
	res, err := SubDateNew(date, substracted.Year(), int(substracted.Month()), substracted.Day())
	if err != nil {
		return YearsMonthsDays{}, err
	}

	return YearsMonthsDays { res.Year(), int(res.Month()), res.Day()}, nil
}

// Calcule une nouvelle date en ajoutant un type Time et un type AnneesMoisHomme
func DatePlusAge(date time.Time, ymd YearsMonthsDays) time.Time {
	return date.AddDate(ymd.Years, ymd.Months, ymd.Days)
}
