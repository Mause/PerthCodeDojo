import time
import unittest
from concurrent.futures import ThreadPoolExecutor

import requests

ROOT_URL = 'http://api.randomuser.me'


class TimeTakenMixin(object):
    class TimeTakenContext(object):
        def __init__(self, test_case, seconds):
            self.test_case = test_case
            self.seconds = seconds

        def __enter__(self):
            self.time_started = time.time()

        def __exit__(self, type, value, traceback):
            passed_seconds = time.time() - self.time_started

            self.test_case.assertLess(
                passed_seconds,
                self.seconds,
                "Should've taken less than {} seconds. Took {} seconds".format(
                    self.seconds,
                    passed_seconds
                )
            )

    def assertTimeTakenLessThan(self, seconds):
        return TimeTakenMixin.TimeTakenContext(self, seconds)


class TestAPI(TimeTakenMixin, unittest.TestCase):
    def test_basic(self):
        self.assertTrue(requests.get(ROOT_URL).json())

    def test_many_individual(self):
        def method(_):
            r = requests.get(ROOT_URL)
            data = r.json()
            self.assertTrue(data)
            print('GG')

            return data

        with self.assertTimeTakenLessThan(80):
            pool = ThreadPoolExecutor(max_workers=10)
            list(pool.map(method, range(50)))

    def test_bulk(self):
        with self.assertTimeTakenLessThan(5):
            r = requests.get(
                ROOT_URL,
                params={
                    'key': '0C4J-SW56-8YOQ-L5ZM',
                    'results': 50
                }
            )

            self.assertEqual(
                len(r.json()['results']),
                50
            )

            self.assertTrue(r.json())


if __name__ == '__main__':
    unittest.main()
