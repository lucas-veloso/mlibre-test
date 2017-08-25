require 'rubygems'
require 'selenium-webdriver'
require 'minitest/autorun'
require './Homepage.rb'
require './Listado.rb'


# Helper class for the test
class Waits
	def initialize(timeout = 10)
		@timeout = timeout
		@wait = Selenium::WebDriver::Wait.new(:timeout => @timeout)
	end

	def wait_until
		@wait.until { yield }
	end
end


# Test using the minitest testing framework for ruby 
class LoginClass < Minitest::Test

	def setup
		@driver = Selenium::WebDriver.for :firefox
		@driver.get "https://mercadolibre.com.ar"
		@driver.manage.window.maximize
	end

	def teardown
		@driver.quit
	end

	# this test searches for iphones on mercadolibre landing page and validates the results
	def test_search_iphone

		# instacing page objects
		home = HomepageObject::Homepage.new(@driver)
		listado = ListadoPageObject::Listadopage.new(@driver)
	
		# search for iphones
		search_text = "Iphone"
		element = home.search_bar()
		assert_equal(element.text, "", "search bar should have no text")
		element.send_keys search_text
		wait = Waits.new

		# validate that the first result mathes my search
		wait.wait_until { home.first_search_result(search_text.downcase).displayed? }
		first_element = home.first_search_result(search_text.downcase)
		assert_equal(first_element.text, search_text.downcase, "text should be the same but downcase")
		assert_equal(first_element.attribute("num"), "1")

		# navigate to the item list
		element.send_keys :arrow_down
		element.send_keys :enter
		search_result = listado.search_name()

		# validate that im searching for iphones
		assert_equal(search_result.text, search_text.downcase, "text should be equal to my search")
		assert_equal(@driver.title, "iPhone en Mercado Libre Argentina", "Message for Iphone")		
	end
end
