# Custom mocks for the pandora repository
mocks :
	mockery --name='(.*)' --case=underscore --dir=redis --output=redis/mocks
	mockery --name='(.*)' --case=underscore --dir=httpclient --output=httpclient/mocks
	mockery --name='(.*)' --case=underscore --dir=onfido --output=onfido/mocks
	mockery --name='(.*)' --case=underscore --dir=digilocker --output=digilocker/mocks
	mockery --name='(.*)' --case=underscore --dir=deposit --output=deposit/mocks
	mockery --name='(.*)' --case=underscore --dir=profile --output=profile/mocks