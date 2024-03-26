# Steps

mkdir portfolio

cd portfolio

python3 -m venv venv

source venv/bin/activate

python3 -m pip install Django

django-admin startproject my_portfolio .

python3 manage.py migrate

python3 manage.py runserver

deactivate

docker image build -t mfk267/django_demo:2024 .

docker run --name django2024 -d -p 8001:8000 mfk267/django_demo:2024

Test in your browser localhost:8001
