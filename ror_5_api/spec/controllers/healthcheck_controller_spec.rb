require 'rails_helper'

RSpec.describe HealthcheckController, type: :controller do

  describe 'GET #index' do
    it 'checks if app is working' do
      get :index
      expect(response.status).to eq(200)
    end
  end


end
