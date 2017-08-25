require 'rubygems'
require 'selenium-webdriver'
require 'minitest/autorun'
require './Homepage.rb'
require './Listado.rb'

class Waits
	def initialize(timeout = 10)
		@timeout = timeout
		@wait = Selenium::WebDriver::Wait.new(:timeout => @timeout)
	end

	def wait_until
		@wait.until { yield }
	end
end



class LoginClass < Minitest::Test

	def setup
		@driver = Selenium::WebDriver.for :firefox
		@driver.get "https://mercadolibre.com.ar"
		@driver.manage.window.maximize
	end

	def teardown
		@driver.quit
	end


	def test_search_iphone
		home = HomepageObject::Homepage.new(@driver)
		listado = ListadoPageObject::Listadopage.new(@driver)
		search_text = "Iphone"
		element = home.search_bar()
		assert_equal(element.text, "", "search bar should have no text")
		element.send_keys search_text
		wait = Waits.new
		wait.wait_until { home.first_search_result(search_text.downcase).displayed? }
		first_element = home.first_search_result(search_text.downcase)
		assert_equal(first_element.text, search_text.downcase, "text should be the same but downcase")
		assert_equal(first_element.attribute("num"), "1")
		element.send_keys :arrow_down
		element.send_keys :enter
		search_result = listado.search_name()
		assert_equal(search_result.text, search_text.downcase, "text should be equal to my search")
		assert_equal(@driver.title, "iPhone en Mercado Libre Argentina", "Message for Iphone")
		
	end
end
