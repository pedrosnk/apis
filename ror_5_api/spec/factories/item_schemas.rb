FactoryGirl.define do
  factory :item_schema do
    properties do
      {'field1' => { 'type' =>'string' } }
    end
    schema 'http://json-schema.org/draft-04/schema#'
    type 'object'
    aditional_properties false
    description ''
    collection_name "genCollection"
    required ['field1']
  end
end
