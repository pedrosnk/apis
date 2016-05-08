FactoryGirl.define do
  factory :item_schema do
    properties do
      {'field1' => { 'type' =>'string' } }
    end
    schema 'http://json-schema.org/draft-04/schema#'
    type 'object'
    aditionalProperties false
    description ''
    collection_name "genCollection"
    required ['field1']

    factory :valid_item_schema do
    end

  end
end
