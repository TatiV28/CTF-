package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/yourusername/mod-ctf/internal/models"
	"github.com/yourusername/mod-ctf/internal/utils"
)

func StartHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"story": "Добро пожаловать в ModNet. ShadowUser распространяет фейки. Начните с OSINT: найдите его email через 'открытый' источник (hint: google 'shadowuser modnet')."})
}

func OSINTHandler(w http.ResponseWriter, r *http.Request) {
	var input struct{ Answer string }
	json.NewDecoder(r.Body).Decode(&input)
	if input.Answer == "shadow@modnet.fake" { // Пример правильного ответа (OSINT)
		json.NewEncoder(w).Encode(models.Stage{Flag: utils.GenerateFlag(1), Next: "Теперь reverse-engineer файл: скачайте /file.bin и найдите скрытый ключ."})
	} else {
		http.Error(w, "Неверно", http.StatusBadRequest)
	}
}

func ReverseHandler(w http.ResponseWriter, r *http.Request) {
	var input struct{ Key string }
	json.NewDecoder(r.Body).Decode(&input)
	if input.Key == "reverse_key_123" { // Проверка reverse-engineering
		json.NewEncoder(w).Encode(models.Stage{Flag: utils.GenerateFlag(2), Next: "Web-sec: exploit SQLi в /profile?id=1 (hint: ' OR 1=1)."})
	} else {
		http.Error(w, "Неверно", http.StatusBadRequest)
	}
}

func WebSecHandler(w http.ResponseWriter, r *http.Request) {
	var input struct{ Exploit string }
	json.NewDecoder(r.Body).Decode(&input)
	if input.Exploit == "' OR 1=1" { // Проверка web-security
		json.NewEncoder(w).Encode(models.Stage{Flag: utils.GenerateFlag(3), Next: "Модерация: примените правило #42 (удалить фейки). Ответ: delete_content."})
	} else {
		http.Error(w, "Неверно", http.StatusBadRequest)
	}
}

func ModerateHandler(w http.ResponseWriter, r *http.Request) {
	var input struct{ Action string }
	json.NewDecoder(r.Body).Decode(&input)
	if input.Action == "delete_content" { // Проверка правил модерации
		json.NewEncoder(w).Encode(models.Stage{Flag: utils.GenerateFlag(4), Next: "CTF завершён! Вы удалили контент по правилам."})
	} else {
		http.Error(w, "Неверно", http.StatusBadRequest)
	}
}
