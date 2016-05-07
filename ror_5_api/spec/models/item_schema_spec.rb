require 'rails_helper'

RSpec.describe ItemSchema, type: :model do

  describe 'validates json schema' do
    it 'with an valid item schema' do
      item_schema = FactoryGirl.build :valid_item_schema
      expect(item_schema).to be_valid
    end
  end

end
