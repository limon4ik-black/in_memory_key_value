package wal

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/limon4ik-black/in_memory_key_value/internal/compute"
	"github.com/limon4ik-black/in_memory_key_value/internal/logger"
)

type WAL struct {
	mutex       sync.RWMutex
	currentFile *os.File
	currentSize int
	maxSize     int
	index       int
	dir         string
	filesWal    []string
}

func InitWal(dir string, maxSize int) (*WAL, error) {

	if err := os.MkdirAll(dir, 0755); err != nil {
		logger.Log.Errorw("failed to create wal directory", "directory", dir, "error", err)
		return nil, err
	}

	wal := &WAL{
		dir:     dir,
		maxSize: maxSize,
		index:   0,
	}

	wal.index, _ = wal.FindLastWALS()
	if wal.index == -1 {
		wal.index = 0
		if err := wal.AddFile(); err != nil {
			return nil, err
		}
	}

	if wal.index != -1 {
		path := fmt.Sprintf("%s/wal_%05d.log", wal.dir, wal.index)
		cur_file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			logger.Log.Errorw("failed to open exist wal-file", "path", path, "error", err)
		}
		wal.currentFile = cur_file
		fileinfo, _ := os.Stat(path)
		wal.currentSize = int(fileinfo.Size())
	}

	return wal, nil
}

func (w *WAL) AddFile() error {

	if w.currentFile != nil {
		w.currentFile.Close()
	}

	path := fmt.Sprintf("%s/wal_%05d.log", w.dir, w.index+1)
	file, err := os.Create(path)
	if err != nil {
		logger.Log.Errorw("failed added new wal-file", "error", err)
		return err
	}
	w.index++
	w.currentFile = file
	w.currentSize = 0

	return nil
}

func (w *WAL) WriteToWal(command string) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	data := []byte(command + "\n")

	if w.currentSize+len(data) > w.maxSize {
		if err := w.AddFile(); err != nil {
			return err
		}
	}

	n, err := w.currentFile.Write(data)
	if err != nil {
		return err
	}
	w.currentSize += n

	if err := w.currentFile.Sync(); err != nil {
		return err
	}

	return nil
}

func (w *WAL) Close() error {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	if w.currentFile != nil {
		return w.currentFile.Close()
	}
	return nil
}

func (w *WAL) FindLastWALS() (int, error) {
	last_index := 0
	path := fmt.Sprintf("%s/wal_*.log", w.dir)
	filesWal, _ := filepath.Glob(path)
	w.filesWal = filesWal
	last_index = len(w.filesWal) - 1
	return last_index, nil
}

func (w *WAL) Load() error {

	for _, file_name := range w.filesWal {
		//path := fmt.Sprintf("%s/%s", w.dir, file_name)
		file, _ := os.Open(file_name)

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			command := scanner.Text()
			_, err := compute.Reception(command)
			if err != nil {
				logger.Log.Errorw("lol")
				return err
			}
		}

		file.Close()
	}

	return nil
}
