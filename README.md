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
    // os
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

Ajouter cette fonction pour afficher une erreur normalisé en cas d'erreur
```go
typotestcolor.AssertNoError(t, err)
```

Cette fonction prend en paramètre :
- `t`: *testing.T, le paramètre des fonctions de tests.
- `err`: une erreur qui à soit la valeur `Error` ou `nil`

Voici un exemple d'utilisation de la fonction typotestcolor.AssertNoError
```go
package tests

import (
    "errors"
    "testing"

    "github.com/typovrak/typotestcolor"
)

func TestTypotestcolorAssertNoError(t *testing.T) {
    t.Run("test typotestcolor.AssertNoError, must return an error", func(t *testing.T) {
        foo := "foo"
        bar := "bar"
        var err error = nil

        if foo != bar {
            err = errors.New("foo does not equal bar")
        }

        typotestcolor.AssertNoError(t, err)
    })

    t.Run("test typotestcolor.AssertNoError, must return nothing", func(t *testing.T) {
        foo := "foo"
        bar := foo
        var err error = nil

        if foo != bar {
            err = errors.New("foo does not equal bar")
        }

        typotestcolor.AssertNoError(t, err)
    })
}
```

Ce code, une fois lancé, retournera le rendu suivant :
_(TODO: mettre le rendu des 2 tests de démonstration)_

Il existe d'autres fonctions d'asserts, fonctionnant sur le même principe que celle-ci avec des objectifs différents. Pour en savoir plus, voir la [documentation technique avancée](#)

### 5. (Bonus) Utilisation de la fonction différentielle pour voir les erreurs dans des strings

Le module contient une fonction différentielle permettant de donner en cas de non équivalence de 2 variables, pouvant avoir des types différents, une erreur normalisé et optimisé pour être comprise en un clin d'oeil
```go
err := typotestcolor.TestDiff(string1, string2, typotestcolor.TestDiffNewDefaultOpts())
```

En cas de non différence, `err == nil`, sinon `err` contient le message d'erreur avec un surlignage voyant de la différence entre les 2 variables.
_(TODO: mettre le rendu des 2 tests de démonstration)_

Voici un exemple d'utilisation de la fonction typotestcolor.TestDiff, version amélioré de l'exemple précedent
```go
package tests

import (
    "errors"
    "testing"

    "github.com/typovrak/typotestcolor"
)

func TestTypotestcolorAssertNoError(t *testing.T) {
    t.Run("test typotestcolor.TestDiff function, must return an error, for demonstration purpose", func(t *testing.T) {
        foo := "foo"
        bar := "bar"
        err := typotestcolor.TestDiff(foo, bar, typotestcolor.TestDiffNewDefaultOpts())

        typotestcolor.AssertNoError(t, err)
    })

    t.Run("test typotestcolor.TestDiff function, must return nothing, for demonstration purpose", func(t *testing.T) {
        foo := "foo"
        bar := foo
        err := typotestcolor.TestDiff(foo, bar, typotestcolor.TestDiffNewDefaultOpts())

        typotestcolor.AssertNoError(t, err)
    })
}
```

Voir le rendu final de ce code:
_(TODO: mettre le rendu des 2 tests de démonstration)_

Avec cette fonction, on vient de remplacer ceci
```go
var err error = nil

if foo != bar {
    err = errors.New("foo does not equal bar")
}
```
par une seule ligne avec une erreur normalisée et possédant un surlignement par défaut rouge vif sur chaque différence entre les 2 variables
```go
err := typotestcolor.TestDiff(foo, bar, typotestcolor.TestDiffNewDefaultOpts())
```

Pour un savoir plus sur les types supportées de cette fonction, voir [cette partie](#) de la documentation technique avancée.














## Customisation

_(TODO: la plus grosse partie de cette documentation)_














## Tests/coverage du module

Voir le détail du coverage du module sur [Codecov](https://app.codecov.io/gh/typovrak/typotestcolor)

Chaque fonction est testé ou est en cours d'évolution sur cette partie afin de garantir un module stable, robuste pour permettre une productivité maximal de vos applications Golang.












## Directives pour l'open source

Dire que j'accepte les issues, avec une description précise + screen ou repo de reproduction minimale. Par contre, je n'accepte pas pour le moment les pull requests, je compte développer moi-même les nouvelles fonctionnalités, améliorations, fix nécessaire à l'évolution du module.

















## Évolutions futures

### Nouvelles fonctionnalités
- [ ] Mettre des couleurs sur la différence de length
- [ ] ajouter une ligne bleu pour afficher la valeur raw, avant transformation en got
- [ ] Ajouter une option pour afficher les tests d'un fichier du X au Y ème
- [ ] Créer une fonction qui me permet de passer mon propre opts dans TestDiff, faire comme pour RunTestColor

### Améliorations et tests
- [ ] Améliorer le makefile afin que le dossier contenant tous les tests soit une variable
- [ ] Tester tous les caractères non ASCII dans chaque fonction (TestDiff, Assert, rendu final)
- [ ] Améliorer les commentaires dans le code
- [ ] log.Fatal ne print aucun résultat, est-ce que cela est corrigeable ou ajouter un paramètre pour annuler les log.Fatal au besoin
- [ ] Mettre printToASCII dans la configuration globale pour simplifier le tout dans l'utilisation de TestDiff
- [ ] Colorier la valeur length et mettre en highlight la différence
- [ ] Tester les groupes de test en Golang, ajouter cela dans le makefile

### Communitation et communauté
- [ ] Valider la documentation auprès de plusieurs développeurs, Golang comme développeur non Golang
- [ ] Faire tous les standards communautaires GitHub
















---

<p align="center"><i>Made with 💜 by <a href="https://typovrak.tv">typovrak</a></i></p>
