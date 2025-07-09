package typotestcolor

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

var DefaultTitle = struct {
	Run         []byte
	Fail        []byte
	Pass        []byte
	Skip        []byte
	Failed      []byte
	Ok          []byte
	ErrorThrown []byte
}{
	Run:         []byte("=== RUN  "),
	Fail:        []byte("--- FAIL:"),
	Pass:        []byte("--- PASS:"),
	Skip:        []byte("--- SKIP:"),
	Failed:      []byte("FAIL"),
	Ok:          []byte("PASS"),
	ErrorThrown: []byte(""),
}

// return exitCode
func RunTestColor(m *testing.M, opts Opts) int {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		// Étape 1 : créer un pipe
		r, w, _ := os.Pipe()

		// Étape 2 : sauvegarder os.Stdout original
		originalStdout := os.Stdout

		// Étape 3 : rediriger os.Stdout vers le writer du pipe
		os.Stdout = w

		// Étape 4 : écrire dans fmt.Println (redirigé)
		fmt.Println("Message capturé")

		// Étape 5 : fermer le writer pour signaler EOF
		w.Close()

		// Étape 6 : restaurer os.Stdout
		os.Stdout = originalStdout

		// Étape 7 : lire tout le contenu capturé depuis le reader
		output, _ := io.ReadAll(r)

		// Étape 8 : afficher ce qui a été capturé
		fmt.Println("Sortie capturée :")
		fmt.Print(string(output))

		os.Exit(1)
	}()

	Debug(opts, "RuntestColor")

	// create a pipe
	r, w, _ := os.Pipe()

	// backup original outputs
	stdout := os.Stdout
	stderr := os.Stderr

	// redirect stdout and stderr to the pipe
	os.Stdout = w
	os.Stderr = w

	// no error when no test executed
	exitCode := 0

	// test mock
	if m != nil {
		exitCode = m.Run()
	}

	// close the writer end of the pipe so the reader stops at EOF
	w.Close()

	// setup the reader
	reader := bufio.NewReader(r)

	errorBefore := false

	// read line by line
	ReadTestLines(opts, reader, stdout, &errorBefore)

	// restore outputs
	os.Stdout = stdout
	os.Stderr = stderr

	// [0, 125]
	return exitCode
}
