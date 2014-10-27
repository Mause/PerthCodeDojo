import requests
import flask
import jinja2

app = flask.Flask(__name__)
TEMPLATE = jinja2.Template(open('template.html').read())

from tests import ROOT_URL


@app.route('/')
def index():
    r = requests.get(ROOT_URL)

    return TEMPLATE.render(r.json()['results'][0])


@app.route('/style.css')
def style():
    return open('style.css').read()

if __name__ == '__main__':
    app.debug = True
    app.run()
