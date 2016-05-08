require 'rails_helper'

RSpec.describe ItemSchema, type: :model do

  describe '#valid?' do
    subject { FactoryGirl.build :item_schema }

    it { expect(subject).to be_valid }
  end

  describe '#to_json_schema' do
    subject { FactoryGirl.build :item_schema }

    let :json_schema do
      { 
        '$schema' => subject[:schema],
        'aditionalProperties' => subject[:aditional_properties],
        'collectionName' => subject[:collection_name],
        'type' => subject[:type],
        'required' => subject[:required],
        'properties' => subject[:properties],
        'description' => subject[:description],
        'id' => subject[:_id].to_s,
      }
    end

    it { expect(subject.to_json_schema).to eql(json_schema) }
  end

end
