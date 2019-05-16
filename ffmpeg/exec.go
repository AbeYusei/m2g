package ffmpeg

import (
	"bufio"
	"log"
	"os/exec"

	"github.com/AbeYusei/m2g"
)

// Exec execute command with stream output
func Exec(dir, c string, args []string) error {

	// コマンドと引数を定義する
	cmd := exec.Command(c, args...)

	// 実行ディレクトリ
	cmd.Dir = dir

	// パイプを作る
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return m2g.Err(m2g.InternalProcess)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return m2g.Err(m2g.InternalProcess)
	}

	err = cmd.Start()
	if err != nil {
		return m2g.Err(m2g.InternalProcess)
	}

	streamReader := func(scanner *bufio.Scanner, outputChan chan string, doneChan chan bool) {
		defer close(outputChan)
		defer close(doneChan)
		for scanner.Scan() {
			outputChan <- scanner.Text()
		}
		doneChan <- true
	}

	// stdout, stderrをひろうgoroutineを起動
	stdoutOutputChan, stdoutDoneChan := make(chan string), make(chan bool)
	stderrOutputChan, stderrDoneChan := make(chan string), make(chan bool)
	go streamReader(bufio.NewScanner(stdout), stdoutOutputChan, stdoutDoneChan)
	go streamReader(bufio.NewScanner(stderr), stderrOutputChan, stderrDoneChan)

	// channel経由でデータを引っこ抜く
	stillGoing := true
	for stillGoing {
		select {
		case <-stdoutDoneChan:
			stillGoing = false
		case line := <-stdoutOutputChan:
			log.Println(line)
		case line := <-stderrOutputChan:
			log.Println(line)
		}
	}

	//一応Waitでプロセスの終了をまつ
	ret := cmd.Wait()
	if ret != nil {
		log.Fatal(err)
	}

	return nil
}
