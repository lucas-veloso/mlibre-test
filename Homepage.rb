module HomepageObject

	class Homepage
		def initialize(webdriver)
			@driver = webdriver
			@url = "https://mercadolibre.com.ar"
		end

		def search_bar()
			return @driver.find_element :class_name => "nav-search-input"
		end

		def first_search_result(item)
				return @driver.find_element :css => "[data-sug~='#{item}']"
		end
	end
end