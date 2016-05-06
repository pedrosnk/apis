class HealthcheckController < ApplicationController

  def index
    redis = Redis.new
    render json: {
      status: "working",
      redis: redis.ping
    }
  end

end
