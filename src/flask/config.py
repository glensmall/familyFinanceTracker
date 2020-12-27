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
    TESTING = False

    # Database
    DATABASE_HOST = environ.get('PROD_DATABASE_HOST')
    DATABASE_USER = environ.get('PROD_DATABASE_USER')
    DATABASE_PASS = environ.get('PROD_DATABASE_PASS')
    DATABASE_NAME = environ.get('PROD_DATABASE_NAME')


# configuration for dev
class DevConfig(Config):

    FLASK_ENV = 'development'
    DEBUG = True
    TESTING = True

    # Database
    DATABASE_HOST = environ.get('DEV_DATABASE_HOST')
    DATABASE_USER = environ.get('DEV_DATABASE_USER')
    DATABASE_PASS = environ.get('DEV_DATABASE_PASS')
    DATABASE_NAME = environ.get('DEV_DATABASE_NAME')

    # TestResponse
    TEST_RESPONSE = 'testData.json'