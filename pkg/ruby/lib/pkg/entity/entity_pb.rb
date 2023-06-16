# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: pkg/entity/entity.proto

require 'google/protobuf'

require 'validate/validate_pb'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("pkg/entity/entity.proto", :syntax => :proto3) do
    add_message "gitlab.agent.entity.AgentMeta" do
      optional :version, :string, 1, json_name: "version"
      optional :commit_id, :string, 2, json_name: "commit_id"
      optional :pod_namespace, :string, 3, json_name: "pod_namespace"
      optional :pod_name, :string, 4, json_name: "pod_name"
    end
    add_message "gitlab.agent.entity.GitalyInfo" do
      optional :address, :string, 1, json_name: "address"
      optional :token, :string, 2, json_name: "token"
      map :features, :string, :string, 3
    end
    add_message "gitlab.agent.entity.GitalyRepository" do
      optional :storage_name, :string, 1, json_name: "storage_name"
      optional :relative_path, :string, 2, json_name: "relative_path"
      optional :gl_repository, :string, 3, json_name: "gl_repository"
      optional :gl_project_path, :string, 4, json_name: "gl_project_path"
    end
  end
end

module Gitlab
  module Agent
    module Entity
      AgentMeta = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("gitlab.agent.entity.AgentMeta").msgclass
      GitalyInfo = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("gitlab.agent.entity.GitalyInfo").msgclass
      GitalyRepository = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("gitlab.agent.entity.GitalyRepository").msgclass
    end
  end
end
