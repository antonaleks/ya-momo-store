package dependencies

import (
	"gitlab.praktikum-services.ru/anton-alekseyev/momo-store/internal/store/dumplings"
	"gitlab.praktikum-services.ru/anton-alekseyev/momo-store/internal/store/dumplings/fake"
)

// NewFakeDumplingsStore returns new fake store for app
func NewFakeDumplingsStore() (dumplings.Store, error) {
	packs := []dumplings.Product{
		{
			ID:          1,
			Name:        "Пельмени",
			Description: "С говядиной",
			Price:       5.00,
			Image:       "https://storage.yandexcloud.net/s3-momo-store-bucket/8dee5a92281746aa887d6f19cf9fdcc7.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEl0dD2qDxM1j6UR-ZPLgj%2F20221015%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20221015T234607Z&X-Amz-Expires=2592000&X-Amz-Signature=DD8EA7D4155218381AF59B59E00753312CAD56933E1F929F17FA9604433F64E9&X-Amz-SignedHeaders=host",
		},
		{
			ID:          2,
			Name:        "Хинкали",
			Description: "Со свининой",
			Price:       3.50,
			Image:       "https://storage.yandexcloud.net/s3-momo-store-bucket/50b583271fa0409fb3d8ffc5872e99bb.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEl0dD2qDxM1j6UR-ZPLgj%2F20221015%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20221015T234715Z&X-Amz-Expires=2592000&X-Amz-Signature=5C73C6640E5CE81809FCDA5EC3AD7B36D3EEA90986F7593713AFA2AC85C16CB4&X-Amz-SignedHeaders=host",
		},
		{
			ID:          3,
			Name:        "Манты",
			Description: "С мясом молодых бычков",
			Price:       2.75,
			Image:       "https://storage.yandexcloud.net/s3-momo-store-bucket/8b50f76f514a4ccaaacdcb832a1b3a2f.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEl0dD2qDxM1j6UR-ZPLgj%2F20221015%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20221015T234744Z&X-Amz-Expires=2592000&X-Amz-Signature=2E6C432A5A53F2090D6D3A51DA7712D969164EAF12C4B1FD299B495A20020857&X-Amz-SignedHeaders=host",
		},
		{
			ID:          4,
			Name:        "Буузы",
			Description: "С телятиной и луком",
			Price:       4.00,
			Image:       "https://storage.yandexcloud.net/s3-momo-store-bucket/788c073d83c14b3fa00675306dfb32b5.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEl0dD2qDxM1j6UR-ZPLgj%2F20221015%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20221015T234744Z&X-Amz-Expires=2592000&X-Amz-Signature=2E6C432A5A53F2090D6D3A51DA7712D969164EAF12C4B1FD299B495A20020857&X-Amz-SignedHeaders=host",
		},
		{
			ID:          5,
			Name:        "Цзяоцзы",
			Description: "С говядиной и свининой",
			Price:       7.25,
			Image:       "https://storage.yandexcloud.net/s3-momo-store-bucket/32cc88a33c3243a6a8838c034878c564.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEl0dD2qDxM1j6UR-ZPLgj%2F20221015%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20221015T234744Z&X-Amz-Expires=2592000&X-Amz-Signature=2E6C432A5A53F2090D6D3A51DA7712D969164EAF12C4B1FD299B495A20020857&X-Amz-SignedHeaders=host",
		},
		{
			ID:          6,
			Name:        "Гедза",
			Description: "С соевым мясом",
			Price:       3.50,
			Image:       "https://storage.yandexcloud.net/s3-momo-store-bucket/7685ad7e9e634a58a4c29120ac5a5ee1.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEl0dD2qDxM1j6UR-ZPLgj%2F20221015%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20221015T234744Z&X-Amz-Expires=2592000&X-Amz-Signature=2E6C432A5A53F2090D6D3A51DA7712D969164EAF12C4B1FD299B495A20020857&X-Amz-SignedHeaders=host",
		},
		{
			ID:          7,
			Name:        "Дим-самы",
			Description: "С уткой",
			Price:       2.65,
			Image:       "https://storage.yandexcloud.net/s3-momo-store-bucket/4bdaeab0ee1842dc888d87d4a435afdd.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEl0dD2qDxM1j6UR-ZPLgj%2F20221015%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20221015T234744Z&X-Amz-Expires=2592000&X-Amz-Signature=2E6C432A5A53F2090D6D3A51DA7712D969164EAF12C4B1FD299B495A20020857&X-Amz-SignedHeaders=host",
		},
		{
			ID:          8,
			Name:        "Момо",
			Description: "С бараниной",
			Price:       5.00,
			Image:       "https://storage.yandexcloud.net/s3-momo-store-bucket/f64dcea998e34278a0006e0a2b104710.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEl0dD2qDxM1j6UR-ZPLgj%2F20221015%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20221015T234744Z&X-Amz-Expires=2592000&X-Amz-Signature=2E6C432A5A53F2090D6D3A51DA7712D969164EAF12C4B1FD299B495A20020857&X-Amz-SignedHeaders=host",
		},
		{
			ID:          9,
			Name:        "Вонтоны",
			Description: "С креветками",
			Price:       4.10,
			Image:       "https://storage.yandexcloud.net/s3-momo-store-bucket/8dee5a92281746aa887d6f19cf9fdcc7.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEl0dD2qDxM1j6UR-ZPLgj%2F20221015%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20221015T234607Z&X-Amz-Expires=2592000&X-Amz-Signature=DD8EA7D4155218381AF59B59E00753312CAD56933E1F929F17FA9604433F64E9&X-Amz-SignedHeaders=host",
		},
		{
			ID:          10,
			Name:        "Баоцзы",
			Description: "С капустой",
			Price:       4.20,
			Image:       "https://storage.yandexcloud.net/s3-momo-store-bucket/50b583271fa0409fb3d8ffc5872e99bb.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEl0dD2qDxM1j6UR-ZPLgj%2F20221015%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20221015T234715Z&X-Amz-Expires=2592000&X-Amz-Signature=5C73C6640E5CE81809FCDA5EC3AD7B36D3EEA90986F7593713AFA2AC85C16CB4&X-Amz-SignedHeaders=host",
		},
		{
			ID:          11,
			Name:        "Кундюмы",
			Description: "С грибами",
			Price:       5.45,
			Image:       "https://storage.yandexcloud.net/s3-momo-store-bucket/8b50f76f514a4ccaaacdcb832a1b3a2f.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEl0dD2qDxM1j6UR-ZPLgj%2F20221015%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20221015T234744Z&X-Amz-Expires=2592000&X-Amz-Signature=2E6C432A5A53F2090D6D3A51DA7712D969164EAF12C4B1FD299B495A20020857&X-Amz-SignedHeaders=host",
		},
		{
			ID:          12,
			Name:        "Курзе",
			Description: "С крабом",
			Price:       3.25,
			Image:       "https://storage.yandexcloud.net/s3-momo-store-bucket/788c073d83c14b3fa00675306dfb32b5.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEl0dD2qDxM1j6UR-ZPLgj%2F20221015%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20221015T234744Z&X-Amz-Expires=2592000&X-Amz-Signature=2E6C432A5A53F2090D6D3A51DA7712D969164EAF12C4B1FD299B495A20020857&X-Amz-SignedHeaders=host",
		},
		{
			ID:          13,
			Name:        "Бораки",
			Description: "С говядиной и бараниной",
			Price:       4.00,
			Image:       "https://storage.yandexcloud.net/s3-momo-store-bucket/7685ad7e9e634a58a4c29120ac5a5ee1.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEl0dD2qDxM1j6UR-ZPLgj%2F20221015%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20221015T234744Z&X-Amz-Expires=2592000&X-Amz-Signature=2E6C432A5A53F2090D6D3A51DA7712D969164EAF12C4B1FD299B495A20020857&X-Amz-SignedHeaders=host",
		},
		{
			ID:          14,
			Name:        "Равиоли",
			Description: "С рикоттой",
			Price:       2.90,
			Image:       "https://storage.yandexcloud.net/s3-momo-store-bucket/4bdaeab0ee1842dc888d87d4a435afdd.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJEl0dD2qDxM1j6UR-ZPLgj%2F20221015%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20221015T234744Z&X-Amz-Expires=2592000&X-Amz-Signature=2E6C432A5A53F2090D6D3A51DA7712D969164EAF12C4B1FD299B495A20020857&X-Amz-SignedHeaders=host",
		},
	}

	store := fake.NewStore()
	store.SetAvailablePacks(packs...)

	return store, nil
}
