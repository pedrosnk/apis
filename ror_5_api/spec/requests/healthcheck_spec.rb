require 'rails_helper'

RSpec.describe 'Healthcheck', type: :request do
  describe 'GET /healthcheck' do

    it 'healthcheck is working ok' do
      get healthcheck_path
      expect(response).to have_http_status(200)
      expect(response.content_type).to eq('application/json')
      expect(response.body).to include('WORKING')
    end

  end
end
