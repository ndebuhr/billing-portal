import unittest
from selenium import webdriver
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.support.ui import Select


class PythonOrgSearch(unittest.TestCase):
    def setUp(self):
        self.driver = webdriver.Remote(
            command_executor="http://selenium:4444/wd/hub",
            desired_capabilities={"browserName": "chrome", "javascriptEnabled": True},
        )

    def test_home(self):
        driver = self.driver
        driver.get("http://static-site")

    def tearDown(self):
        self.driver.close()


if __name__ == "__main__":
    unittest.main()
