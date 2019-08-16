import time
import os

from selenium.common.exceptions import WebDriverException
from selenium.webdriver.chrome.options import Options
from selenium import webdriver
import dotenv


# Load settings from environment
dotenv.load_dotenv()


class Robot:
    """
    Robot.

    Selenium robot class
    """
    def __init__(self):

        # Settings
        self.selenium_url = os.environ.get('SELENIUM_URL', 'http://localhost:4444/wd/hub')
        self.headless = os.environ.get('HEADLESS', None) is not None
        self.user_agent = os.environ.get('USER_AGENT', 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36')
        self.locale = os.environ.get('LOCALE', 'ru,ru_RU')
        self.chromedriver_path = os.environ.get('CHROMEDRIVER_PATH', '/usr/local/bin/chromedriver')
        # TODO: other robot settings
        # self.param = os.environ.get('PARAM', 'default_value')

        # Get driver
        self.driver = self.get_driver()

    def start(self):
        """
        Start

        Exec job
        :return:
        """
        # TODO write code here :)

    def get_driver(self):
        """
        Get Driver

        Initialize webdriver.
        :return:
        """
        opts = self.build_options()

        if self.in_docker():
            return webdriver.Remote(command_executor=self.selenium_url, desired_capabilities=opts.to_capabilities())
        else:
            return webdriver.Chrome(self.chromedriver_path, chrome_options=opts)

    def build_options(self):
        """
        Build Options

        Build chrome driver options
        :return:
        """
        opts = Options()
        opts.headless = self.headless
        opts.add_argument('user-agent=%s' % self.user_agent)
        opts.add_argument('start-maximized')
        opts.add_argument('disable-infobars')
        opts.add_experimental_option("useAutomationExtension", False)
        opts.add_experimental_option("forceDevToolsScreenshot", True)
        opts.add_experimental_option("excludeSwitches", ['enable-automation'])
        opts.add_experimental_option("prefs", {
            'intl.accept_languages': self.locale,
            'profile.default_content_settings.popups': 0,
        })
        opts.set_capability('goog:loggingPrefs', {'performance': 'ALL'})

        return opts

    @staticmethod
    def in_docker():
        """
        In Docker

        Check robot ran in docker
        :return:
        """
        return os.path.exists('/.dockerenv')


if __name__ == '__main__':
    robot = Robot()
    robot.start()
