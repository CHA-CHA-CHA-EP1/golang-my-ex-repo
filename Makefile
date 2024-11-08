repo:
	mockgen -source=domain/repositories/user.go -destination=mocks/mock_user_repository.go -package=mocks

service:
	mockgen -source=domain/services/user.go -destination=mocks/mock_user_service.go -package=mocks

