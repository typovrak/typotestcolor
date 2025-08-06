[![Go 1.23.6+](https://img.shields.io/badge/Go-1.23.6%2B-a6e3a1?labelColor=45475a)](https://go.dev)
[![License MIT](https://img.shields.io/badge/License-MIT-cba6f7.svg?labelColor=45475a)](LICENSE.md)
[![codecov](https://codecov.io/gh/typovrak/typotestcolor/graph/badge.svg?token=L1MPWRJM6O)](https://codecov.io/gh/typovrak/typotestcolor)
[![Buy me a coffee](https://img.shields.io/badge/Buy%20me%20a%20coffee-☕-fab387?labelColor=45475a)](https://typovrak.tv/coffee)
[![Discord join us](https://img.shields.io/badge/Discord-Join%20us-74c7ec?labelColor=45475a&logo=discord&logoColor=white)](https://typovrak.tv/discord)







# Typotestcolor est un module Golang sans dépendances afin de rendre vos tests lisibles et compréhensible.

Typotestcolor est un module Golang permettant d'avoir un rendu des tests colorés, compréhensible, 100% customisable avec des fonctionnalités uniques tel que l'aggrégation de tests ainsi qu'un résumé de tous les tests validés, ratés ou skippés.








## Typotestcolor en action sur un projet avec + de 200 tests

### Rendu des tests avec des logs d'erreurs
<img src="/assets/typotestcolor-run-min.jpg" alt="">

### Résultat des tests avec résumé compact
<img src="/assets/typotestcolor-result-min.jpg" alt="">

### Rendu en direct!
_TODO: vidéo du rendu avec asciinema.org_








## Liste des fonctionnalités

### 1. Résumé des tests

### 2. Aggrégation de groupe de ligne

### 3. Customisation des préfix, suffixes de chaque ligne, type de ligne ainsi que d'en-tête de groupe de ligne

### 4. Chaque couleur est modifiable à volonté avec des codes ANSI. Des helpers sont présent pour que vous n'ayez aucune chose à faire à part atteindre votre objectif de coverage.

### 5. Fonction différentielle pour trouver l'erreur dans une string en un clin d'oeil

### 6. Aucune perte de performance sur les tests, seulement du gain de temps pour vous
- utilisation d'une goroutine pour avoir un rendu le plus rapide possible!

### 7. Module robuste testé dans son intégralité avec coverage et CI
- chaque fonction est testé 

### 8. Fonctions d'asserts pour valider vos tests en quelques lignes

### 9. Fonctionne out-of-the-box en une seule ligne!








## Utilisation

_TODO: 1. vidéo de l'installation du module + mise en place dans main_test_
_TODO: 2. vidéo de la mise en place des asserts et fonctions différentielles avec asciinema.org_

### 1. Ajout du module dans votre application Golang

Dans le fichier `go.mod`, ajouter cette ligne avec la version souhaitée (ici, la plus récente stable)
```go
require github.com/typovrak/typotestcolor v1.1.0
```

### 2. Utilisation du module avec la configuration par défaut pour lancer les tests

Dans le fichier `main_test.go` ou votre fichier qui contient la ligne `m.Run()`
```go
package tests

import (
    // "os"
    "testing"

    "github.com/typovrak/typotestcolor"
)

// INFO: utile pour utiliser la configuration du module typotestcolor dans les tests plus tard
// cela permet de garder un rendu homogène si vous customisez des couleurs, préfix, etc
var Opts = typotestcolor.NewDefaultOpts()

func TestMain(m *testing.M) {
    // INFO: mettre vos variables d'environnements pour les tests ici
    // os.Setenv("APP_GO_TEST", "true")
	
    // INFO: remplacer l'exécution par défaut par le module typotestcolor
    // exitCode := m.Run()

    exitCode := typotestcolor.RunTestColor(m, Opts)
    os.Exit(exitCode)
}

```

Maintenant, il ne vous reste plus qu'à lancer vos tests pour avoir le même rendu que sur les images et vidéos de ce [README](#), qui est ma configuration quotidienne !

Si vous souhaitez customiser votre rendu par des couleurs, préfix, suffix, header, footer, modifier le résumé, etc, ceci est expliqué [ici](#)

### 3. (Bonus) Lancer les tests avec un makefile (TODO: pertinant?)

Voici un `makefile` par défaut si tous vos tests sont dans un même dossier, ici `tests`
```makefile
CURRENT_DATE=$(shell date +%Y-%m-%dT%H:%M:%S%z)
COVERAGE_FILE=./coverage.txt

.PHONY: test
test:
    @go test ./tests/... -v

race:
    @go test -race ./tests/... -v

.PHONY: coverage
coverage:
    @go test ./tests/... -v -coverprofile=$(COVERAGE_FILE) -coverpkg=./...

.PHONY: show-coverage
show-coverage:
    @go tool cover -html=$(COVERAGE_FILE)

.PHONY: fmt
fmt:
    @gofmt -l .
```

Ce makefile permet à la fois de :
- lancer tous les tests : `make test`
- lancer tous les tests en vérifiant qu'aucun race conditions n'est présent : `make race`
- lancer tous les tests et créer un rapport de coverage : `make coverage`
- visualiser le résultat du coverage dans un navigateur web : `make show-coverage`
- formatter tous les fichiers de l'application Golang avec le formatteur recommandé gofmt : `make fmt`

Il suffit donc d'exécuter la commande `make test` ou `go test ./tests/... -v` afin d'avoir un rendu moderne et optimisé de vos tests !

### 4. (Bonus) Mise en place d'asserts pour valider les tests

### 5. (Bonus) Utilisation de la fonction différentielle pour voir les erreurs dans des strings









## Customisation

## Tests/coverage du module

## Directives pour l'open source

## Évolutions futures





















---

<p align="center"><i>Made with 💜 by <a href="https://typovrak.tv">typovrak</a></i></p>



