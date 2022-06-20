# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: internal/module/configuration_project/rpc/rpc.proto

require 'google/protobuf'

require 'validate/validate_pb'
require 'internal/module/modserver/modserver_pb'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("internal/module/configuration_project/rpc/rpc.proto", :syntax => :proto3) do
    add_message "gitlab.agent.configuration_project.rpc.ListAgentConfigFilesRequest" do
      optional :repository, :message, 1, "gitlab.agent.modserver.Repository", json_name: "repository"
      optional :gitaly_address, :message, 2, "gitlab.agent.modserver.GitalyAddress", json_name: "gitalyAddress"
      optional :default_branch, :string, 3, json_name: "defaultBranch"
    end
    add_message "gitlab.agent.configuration_project.rpc.ListAgentConfigFilesResponse" do
      repeated :config_files, :message, 1, "gitlab.agent.configuration_project.rpc.AgentConfigFile", json_name: "config_files"
    end
    add_message "gitlab.agent.configuration_project.rpc.AgentConfigFile" do
      optional :name, :string, 1, json_name: "name"
      optional :agent_name, :string, 2, json_name: "agent_name"
    end
  end
end

module Gitlab
  module Agent
    module ConfigurationProject
      module Rpc
        ListAgentConfigFilesRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("gitlab.agent.configuration_project.rpc.ListAgentConfigFilesRequest").msgclass
        ListAgentConfigFilesResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("gitlab.agent.configuration_project.rpc.ListAgentConfigFilesResponse").msgclass
        AgentConfigFile = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("gitlab.agent.configuration_project.rpc.AgentConfigFile").msgclass
      end
    end
  end
end
