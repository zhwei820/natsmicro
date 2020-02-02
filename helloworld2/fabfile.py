# coding=utf-8

from fabric.api import (
    local,
    with_settings,
    hosts,
    cd,
    run,
    execute,
    env,
)
from fabric.context_managers import shell_env


PRO_ENV = '10.9.253.65'
PRO_USER = 'data'

TEST_ENV = '10.9.255.43'
TEST_USER = 'data'

env.user = 'data'


def runserver():
    '''
    start dev django server
    '''
    local(
        'gf run main.go'
    )

def djangoserver():
    '''
    start dev django server
    '''
    local(
        'cd py/django_migrate; python3 manage.py runserver 0.0.0.0:8000'
    )


def pipinstall():
    '''
    pip install
    '''
    local('cd py/django_migrate; pip3 install -r requirments.txt')


def manage(cmd):
    '''
    run django manage command
    '''
    local('cd py/django_migrate; python3 manage.py {cmd}'.format(cmd=cmd))


def makemigrations(merge=None):
    '''
    database makemigrations
    '''
    cmd = 'makemigrations' if merge is None else 'makemigrations --merge'
    manage(cmd)


def migrate(ENV='dev'):
    '''
    database migrate
    '''
    with shell_env(API_ENV=ENV):
        manage('migrate')


def showmigrations(args=''):
    '''
    show migrations
    '''
    manage('%s %s' % ('showmigrations', args))


def shell():
    '''
    shell
    '''
    manage('shell')


def create_db(db='hello1'):
    '''
    create db if db not exists
    '''
    local(
        'docker exec -it product_mysql mysql -uroot -prootroot -e '
        '"CREATE DATABASE IF NOT EXISTS {db} DEFAULT CHARSET utf8 COLLATE utf8_general_ci;"'.format(db=db)
    )


# for deploy_test_from_local
def server_db_migrate(container='pro'):
    with cd('/data/opt/contentapi'):
        run('docker-compose pull web')
        run('docker-compose run --rm %s python3 manage.py migrate' % container)


def server_service_restart(container='pro'):
    with cd('/data/opt/contentapi'):
        run('docker-compose pull web')
        run('docker-compose stop %s' % container)
        run('docker-compose rm -f %s' % container)
        run('docker-compose up -d %s' % container)


def update_image():
    local('rm -rf ./bin')
    local('gf build main.go -o ./bin/linux_amd64/main')
    local('docker build . -t zw/hello:latest')


def update_models():
    local('gf gen model ./app/model -l "mysql:root:rootroot@tcp(127.0.0.1:3306)/hello1" -y')
    local('rm ./app/model/auth_* -rf;')
    local('rm ./app/model/django* -rf;')
