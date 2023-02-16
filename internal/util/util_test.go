package util

import (
	"os"
	"testing"
	"time"
)

func TestDirExists(t *testing.T) {
	dirPath := "./testDir"

	// Создаем тестовую директорию
	if err := os.Mkdir(dirPath, os.ModePerm); err != nil {
		t.Fatalf("Не удалось создать тестовую директорию: %s", err)
	}
	defer os.Remove(dirPath) // Удаляем тестовую директорию по завершению

	// Проверяем существование директории
	exists, err := DirExists(dirPath)
	if err != nil {
		t.Fatalf("Ошибка при вызове DirExists: %s", err)
	}
	if !exists {
		t.Fatalf("DirExists вернула false, ожидалось true")
	}

	// Проверяем несуществование директории
	exists, err = DirExists(dirPath + "_not_exist")
	if err != nil {
		t.Fatalf("Ошибка при вызове DirExists для несуществующей директории: %s", err)
	}
	if exists {
		t.Fatalf("DirExists вернула true для несуществующей директории, ожидалось false")
	}
}

func TestMakeDir(t *testing.T) {
	dirPath := "./testMakeDir"

	// Создаем директорию
	if err := MakeDir(dirPath); err != nil {
		t.Fatalf("Не удалось создать директорию: %s", err)
	}
	defer os.Remove(dirPath) // Удаляем директорию по завершению

	// Проверяем, что директория существует
	exists, err := DirExists(dirPath)
	if err != nil {
		t.Fatalf("Ошибка при проверке существования директории: %s", err)
	}
	if !exists {
		t.Fatalf("Директория не создана")
	}
}

func TestGetMD5Hash(t *testing.T) {
	text := "Hello, world!"
	expectedHash := "6cd3556deb0da54bca060b4c39479839" // известный MD5-хеш для строки "Hello, world!"
	hash := GetMD5Hash(text)
	if hash != expectedHash {
		t.Fatalf("GetMD5Hash вернула неправильный хеш: получено %s, ожидалось %s", hash, expectedHash)
	}
}

func TestCreateExpireDate(t *testing.T) {
	now := time.Now()
	expireTime := 5 * time.Minute
	expireDate := CreateExpireDate(expireTime)

	// Проверяем, что expireDate корректно установлено
	if (expireDate.Sub(now)/time.Second)*time.Second != expireTime {
		t.Fatalf("CreateExpireDate вернула неправильную дату истечения срока")
	}
}

func TestCompareDate(t *testing.T) {
	now := time.Now()
	future := now.Add(5 * time.Minute)
	past := now.Add(-5 * time.Minute)

	if !CompareDate(now, future) {
		t.Fatal("CompareDate вернула false для даты в будущем")
	}

	if !CompareDate(now, now) {
		t.Fatal("CompareDate вернула false для одинаковых дат")
	}

	if CompareDate(future, past) {
		t.Fatal("CompareDate вернула true для даты в прошлом")
	}
}
