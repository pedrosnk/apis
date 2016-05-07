require 'rails_helper'

RSpec.describe "ItemSchemas", type: :request do
  describe "GET /item_schemas" do
    it "works! (now write some real specs)" do
      get item_schemas_path
      expect(response).to have_http_status(200)
    end
  end
end
