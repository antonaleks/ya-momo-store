package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.praktikum-services.ru/anton-alekseyev/momo-store/cmd/api/dependencies"
)

func TestFakeAppIntegrational(t *testing.T) {
	store, err := dependencies.NewFakeDumplingsStore()
	assert.NoError(t, err)
	app, err := NewInstance(store)
	assert.NoError(t, err)

	t.Run("create_order", func(t *testing.T) {
		for i := 1; i <= 10; i++ {
			t.Run("id"+strconv.Itoa(i), func(t *testing.T) {
				r := httptest.NewRequest("POST", "/orders", nil)
				w := httptest.NewRecorder()
				app.CreateOrderController(w, r)

				assert.Equal(t, http.StatusOK, w.Code)
				assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
				fmt.Fprintln(os.Stdout, "_____")
				fmt.Fprintln(os.Stdout, w.Body.String())
				fmt.Fprintln(os.Stdout, "_____")

				expectedJSON, err := json.Marshal(map[string]interface{}{"id": i})
				assert.NoError(t, err)
				assert.JSONEq(t, string(expectedJSON), w.Body.String())
			})
		}
	})

	t.Run("list_dumplings", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/packs", nil)
		w := httptest.NewRecorder()
		app.ListDumplingsController(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

		fmt.Fprintln(os.Stdout, "_____")
		fmt.Fprintln(os.Stdout, w.Body.String())
		fmt.Fprintln(os.Stdout, "_____")

		expectedJSON := "{\"results\":[{\"id\":1,\"name\":\"Пельмени\",\"price\":5,\"description\":\"С говядиной\",\"image\":\"https://storage.yandexcloud.net/s3-momo-store-bucket/8dee5a92281746aa887d6f19cf9fdcc7.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEl0dD2qDxM1j6UR-ZPLgj%2F20221015%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20221015T234607Z&X-Amz-Expires=2592000&X-Amz-Signature=DD8EA7D4155218381AF59B59E00753312CAD56933E1F929F17FA9604433F64E9&X-Amz-SignedHeaders=host\"},{\"id\":2,\"name\":\"Хинкали\",\"price\":3.5,\"description\":\"Со свининой\",\"image\":\"https://storage.yandexcloud.net/s3-momo-store-bucket/50b583271fa0409fb3d8ffc5872e99bb.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEl0dD2qDxM1j6UR-ZPLgj%2F20221015%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20221015T234715Z&X-Amz-Expires=2592000&X-Amz-Signature=5C73C6640E5CE81809FCDA5EC3AD7B36D3EEA90986F7593713AFA2AC85C16CB4&X-Amz-SignedHeaders=host\"},{\"id\":3,\"name\":\"Манты\",\"price\":2.75,\"description\":\"С мясом молодых бычков\",\"image\":\"https://storage.yandexcloud.net/s3-momo-store-bucket/8b50f76f514a4ccaaacdcb832a1b3a2f.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEl0dD2qDxM1j6UR-ZPLgj%2F20221015%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20221015T234744Z&X-Amz-Expires=2592000&X-Amz-Signature=2E6C432A5A53F2090D6D3A51DA7712D969164EAF12C4B1FD299B495A20020857&X-Amz-SignedHeaders=host\"},{\"id\":4,\"name\":\"Буузы\",\"price\":4,\"description\":\"С телятиной и луком\",\"image\":\"https://storage.yandexcloud.net/s3-momo-store-bucket/788c073d83c14b3fa00675306dfb32b5.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEl0dD2qDxM1j6UR-ZPLgj%2F20221015%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20221015T234744Z&X-Amz-Expires=2592000&X-Amz-Signature=2E6C432A5A53F2090D6D3A51DA7712D969164EAF12C4B1FD299B495A20020857&X-Amz-SignedHeaders=host\"},{\"id\":5,\"name\":\"Цзяоцзы\",\"price\":7.25,\"description\":\"С говядиной и свининой\",\"image\":\"https://storage.yandexcloud.net/s3-momo-store-bucket/32cc88a33c3243a6a8838c034878c564.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEl0dD2qDxM1j6UR-ZPLgj%2F20221015%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20221015T234744Z&X-Amz-Expires=2592000&X-Amz-Signature=2E6C432A5A53F2090D6D3A51DA7712D969164EAF12C4B1FD299B495A20020857&X-Amz-SignedHeaders=host\"},{\"id\":6,\"name\":\"Гедза\",\"price\":3.5,\"description\":\"С соевым мясом\",\"image\":\"https://storage.yandexcloud.net/s3-momo-store-bucket/7685ad7e9e634a58a4c29120ac5a5ee1.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEl0dD2qDxM1j6UR-ZPLgj%2F20221015%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20221015T234744Z&X-Amz-Expires=2592000&X-Amz-Signature=2E6C432A5A53F2090D6D3A51DA7712D969164EAF12C4B1FD299B495A20020857&X-Amz-SignedHeaders=host\"},{\"id\":7,\"name\":\"Дим-самы\",\"price\":2.65,\"description\":\"С уткой\",\"image\":\"https://storage.yandexcloud.net/s3-momo-store-bucket/4bdaeab0ee1842dc888d87d4a435afdd.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEl0dD2qDxM1j6UR-ZPLgj%2F20221015%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20221015T234744Z&X-Amz-Expires=2592000&X-Amz-Signature=2E6C432A5A53F2090D6D3A51DA7712D969164EAF12C4B1FD299B495A20020857&X-Amz-SignedHeaders=host\"},{\"id\":8,\"name\":\"Момо\",\"price\":5,\"description\":\"С бараниной\",\"image\":\"https://storage.yandexcloud.net/s3-momo-store-bucket/f64dcea998e34278a0006e0a2b104710.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEl0dD2qDxM1j6UR-ZPLgj%2F20221015%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20221015T234744Z&X-Amz-Expires=2592000&X-Amz-Signature=2E6C432A5A53F2090D6D3A51DA7712D969164EAF12C4B1FD299B495A20020857&X-Amz-SignedHeaders=host\"},{\"id\":9,\"name\":\"Вонтоны\",\"price\":4.1,\"description\":\"С креветками\",\"image\":\"https://storage.yandexcloud.net/s3-momo-store-bucket/8dee5a92281746aa887d6f19cf9fdcc7.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEl0dD2qDxM1j6UR-ZPLgj%2F20221015%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20221015T234744Z&X-Amz-Expires=2592000&X-Amz-Signature=2E6C432A5A53F2090D6D3A51DA7712D969164EAF12C4B1FD299B495A20020857&X-Amz-SignedHeaders=host?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEl0dD2qDxM1j6UR-ZPLgj%2F20221015%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20221015T234607Z&X-Amz-Expires=2592000&X-Amz-Signature=DD8EA7D4155218381AF59B59E00753312CAD56933E1F929F17FA9604433F64E9&X-Amz-SignedHeaders=host\"},{\"id\":10,\"name\":\"Баоцзы\",\"price\":4.2,\"description\":\"С капустой\",\"image\":\"https://storage.yandexcloud.net/s3-momo-store-bucket/50b583271fa0409fb3d8ffc5872e99bb.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEl0dD2qDxM1j6UR-ZPLgj%2F20221015%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20221015T234744Z&X-Amz-Expires=2592000&X-Amz-Signature=2E6C432A5A53F2090D6D3A51DA7712D969164EAF12C4B1FD299B495A20020857&X-Amz-SignedHeaders=host?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEl0dD2qDxM1j6UR-ZPLgj%2F20221015%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20221015T234715Z&X-Amz-Expires=2592000&X-Amz-Signature=5C73C6640E5CE81809FCDA5EC3AD7B36D3EEA90986F7593713AFA2AC85C16CB4&X-Amz-SignedHeaders=host\"},{\"id\":11,\"name\":\"Кундюмы\",\"price\":5.45,\"description\":\"С грибами\",\"image\":\"https://storage.yandexcloud.net/s3-momo-store-bucket/8b50f76f514a4ccaaacdcb832a1b3a2f.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEl0dD2qDxM1j6UR-ZPLgj%2F20221015%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20221015T234744Z&X-Amz-Expires=2592000&X-Amz-Signature=2E6C432A5A53F2090D6D3A51DA7712D969164EAF12C4B1FD299B495A20020857&X-Amz-SignedHeaders=host?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEl0dD2qDxM1j6UR-ZPLgj%2F20221015%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20221015T234744Z&X-Amz-Expires=2592000&X-Amz-Signature=2E6C432A5A53F2090D6D3A51DA7712D969164EAF12C4B1FD299B495A20020857&X-Amz-SignedHeaders=host\"},{\"id\":12,\"name\":\"Курзе\",\"price\":3.25,\"description\":\"С крабом\",\"image\":\"https://storage.yandexcloud.net/s3-momo-store-bucket/788c073d83c14b3fa00675306dfb32b5.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEl0dD2qDxM1j6UR-ZPLgj%2F20221015%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20221015T234744Z&X-Amz-Expires=2592000&X-Amz-Signature=2E6C432A5A53F2090D6D3A51DA7712D969164EAF12C4B1FD299B495A20020857&X-Amz-SignedHeaders=host\"},{\"id\":13,\"name\":\"Бораки\",\"price\":4,\"description\":\"С говядиной и бараниной\",\"image\":\"https://storage.yandexcloud.net/s3-momo-store-bucket/7685ad7e9e634a58a4c29120ac5a5ee1.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEl0dD2qDxM1j6UR-ZPLgj%2F20221015%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20221015T234744Z&X-Amz-Expires=2592000&X-Amz-Signature=2E6C432A5A53F2090D6D3A51DA7712D969164EAF12C4B1FD299B495A20020857&X-Amz-SignedHeaders=host\"},{\"id\":14,\"name\":\"Равиоли\",\"price\":2.9,\"description\":\"С рикоттой\",\"image\":\"https://storage.yandexcloud.net/s3-momo-store-bucket/4bdaeab0ee1842dc888d87d4a435afdd.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEl0dD2qDxM1j6UR-ZPLgj%2F20221015%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20221015T234744Z&X-Amz-Expires=2592000&X-Amz-Signature=2E6C432A5A53F2090D6D3A51DA7712D969164EAF12C4B1FD299B495A20020857&X-Amz-SignedHeaders=host\"}]}\n"

		assert.NoError(t, err)
		assert.JSONEq(t, string(expectedJSON), w.Body.String())
	})

	t.Run("healthcheck", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/health", nil)
		w := httptest.NewRecorder()
		app.HealthcheckController(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}
