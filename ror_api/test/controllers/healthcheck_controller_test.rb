require 'test_helper'

class HealthcheckControllerTest < ActionController::TestCase
  test "should get index" do
    get :index
    assert_response :success
  end

end
