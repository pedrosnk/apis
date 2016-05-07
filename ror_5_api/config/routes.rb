Rails.application.routes.draw do
  resources :item_schemas
  get 'healthcheck' => 'healthcheck#index'

  # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html
end
