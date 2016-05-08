require 'rails_helper'

RSpec.describe ItemSchema, type: :model do

  describe '#valid?' do
    subject { FactoryGirl.build :item_schema }

    it { expect(subject).to be_valid }
  end

end
