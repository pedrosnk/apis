class HealthcheckController < ApplicationController
  def index
    render json: { status: 'WORKING' }
  end
end
