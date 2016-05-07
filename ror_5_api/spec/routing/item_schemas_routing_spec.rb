require "rails_helper"

RSpec.describe ItemSchemasController, type: :routing do
  describe "routing" do

    it "routes to #index" do
      expect(:get => "/item_schemas").to route_to("item_schemas#index")
    end

    it "routes to #show" do
      expect(:get => "/item_schemas/1").to route_to("item_schemas#show", :id => "1")
    end

    it "routes to #create" do
      expect(:post => "/item_schemas").to route_to("item_schemas#create")
    end

    it "routes to #update via PUT" do
      expect(:put => "/item_schemas/1").to route_to("item_schemas#update", :id => "1")
    end

    it "routes to #update via PATCH" do
      expect(:patch => "/item_schemas/1").to route_to("item_schemas#update", :id => "1")
    end

    it "routes to #destroy" do
      expect(:delete => "/item_schemas/1").to route_to("item_schemas#destroy", :id => "1")
    end

  end
end
