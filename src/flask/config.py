from os import environ, path
from dotenv import load_dotenv

basedir = path.abspath(path.dirname(__file__))
load_dotenv(path.join(basedir, '.env'))

# base configuration class
class Config:

    SECRET_KEY = environ.get('SECRET_KEY')
    STATIC_FOLDER = 'static'
    TEMPLATES_FOLDER = 'templates'
    SESSION_COOKIE_NAME = environ.get('SESSION_COOKIE_NAME')

    



# configuration for production
class ProdConfig(Config):

    FLASK_ENV = 'production'
    DEBUG = False
    Testing = False

    # Database
    DatabaseHost = environ.get('PROD_DATABASE_HOST')
    DatabaseUser = environ.get('PROD_DATABASE_USER')
    DatabasePass = environ.get('PROD_DATABASE_PASS')
    DatabaseName = environ.get('PROD_DATABASE_NAME')


# configuration for dev
class DevConfig(Config):

    FLASK_ENV = 'development'
    DEBUG = True
    Testing - True

    # Database
    DatabaseHost = environ.get('DEV_DATABASE_HOST')
    DatabaseUser = environ.get('DEV_DATABASE_USER')
    DatabasePass = environ.get('DEV_DATABASE_PASS')
    DatabaseName = environ.get('DEV_DATABASE_NAME')