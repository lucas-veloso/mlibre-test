module ListadoPageObject
	class Listadopage
		def initialize(webdriver)
			@driver = webdriver
			@url = "https://listado.mercadolibre.com.ar"
		end

		def search_name()
			return @driver.find_element :class_name => "breadcrumb__title"
		end
	end
end
