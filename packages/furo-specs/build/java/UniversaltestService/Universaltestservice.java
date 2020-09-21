// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: UniversaltestService/universaltestservice.proto

package UniversaltestService;

public final class Universaltestservice {
  private Universaltestservice() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  public interface CreateUniversaltestServiceRequestOrBuilder extends
      // @@protoc_insertion_point(interface_extends:UniversaltestService.CreateUniversaltestServiceRequest)
      com.google.protobuf.MessageOrBuilder {

    /**
     * <code>.universaltest.Universaltest data = 1;</code>
     */
    boolean hasData();
    /**
     * <code>.universaltest.Universaltest data = 1;</code>
     */
    universaltest.UniversaltestOuterClass.Universaltest getData();
    /**
     * <code>.universaltest.Universaltest data = 1;</code>
     */
    universaltest.UniversaltestOuterClass.UniversaltestOrBuilder getDataOrBuilder();
  }
  /**
   * Protobuf type {@code UniversaltestService.CreateUniversaltestServiceRequest}
   */
  public  static final class CreateUniversaltestServiceRequest extends
      com.google.protobuf.GeneratedMessageV3 implements
      // @@protoc_insertion_point(message_implements:UniversaltestService.CreateUniversaltestServiceRequest)
      CreateUniversaltestServiceRequestOrBuilder {
  private static final long serialVersionUID = 0L;
    // Use CreateUniversaltestServiceRequest.newBuilder() to construct.
    private CreateUniversaltestServiceRequest(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
      super(builder);
    }
    private CreateUniversaltestServiceRequest() {
    }

    @java.lang.Override
    public final com.google.protobuf.UnknownFieldSet
    getUnknownFields() {
      return this.unknownFields;
    }
    private CreateUniversaltestServiceRequest(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      this();
      if (extensionRegistry == null) {
        throw new java.lang.NullPointerException();
      }
      int mutable_bitField0_ = 0;
      com.google.protobuf.UnknownFieldSet.Builder unknownFields =
          com.google.protobuf.UnknownFieldSet.newBuilder();
      try {
        boolean done = false;
        while (!done) {
          int tag = input.readTag();
          switch (tag) {
            case 0:
              done = true;
              break;
            case 10: {
              universaltest.UniversaltestOuterClass.Universaltest.Builder subBuilder = null;
              if (data_ != null) {
                subBuilder = data_.toBuilder();
              }
              data_ = input.readMessage(universaltest.UniversaltestOuterClass.Universaltest.parser(), extensionRegistry);
              if (subBuilder != null) {
                subBuilder.mergeFrom(data_);
                data_ = subBuilder.buildPartial();
              }

              break;
            }
            default: {
              if (!parseUnknownField(
                  input, unknownFields, extensionRegistry, tag)) {
                done = true;
              }
              break;
            }
          }
        }
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.setUnfinishedMessage(this);
      } catch (java.io.IOException e) {
        throw new com.google.protobuf.InvalidProtocolBufferException(
            e).setUnfinishedMessage(this);
      } finally {
        this.unknownFields = unknownFields.build();
        makeExtensionsImmutable();
      }
    }
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return UniversaltestService.Universaltestservice.internal_static_UniversaltestService_CreateUniversaltestServiceRequest_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return UniversaltestService.Universaltestservice.internal_static_UniversaltestService_CreateUniversaltestServiceRequest_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest.class, UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest.Builder.class);
    }

    public static final int DATA_FIELD_NUMBER = 1;
    private universaltest.UniversaltestOuterClass.Universaltest data_;
    /**
     * <code>.universaltest.Universaltest data = 1;</code>
     */
    public boolean hasData() {
      return data_ != null;
    }
    /**
     * <code>.universaltest.Universaltest data = 1;</code>
     */
    public universaltest.UniversaltestOuterClass.Universaltest getData() {
      return data_ == null ? universaltest.UniversaltestOuterClass.Universaltest.getDefaultInstance() : data_;
    }
    /**
     * <code>.universaltest.Universaltest data = 1;</code>
     */
    public universaltest.UniversaltestOuterClass.UniversaltestOrBuilder getDataOrBuilder() {
      return getData();
    }

    private byte memoizedIsInitialized = -1;
    @java.lang.Override
    public final boolean isInitialized() {
      byte isInitialized = memoizedIsInitialized;
      if (isInitialized == 1) return true;
      if (isInitialized == 0) return false;

      memoizedIsInitialized = 1;
      return true;
    }

    @java.lang.Override
    public void writeTo(com.google.protobuf.CodedOutputStream output)
                        throws java.io.IOException {
      if (data_ != null) {
        output.writeMessage(1, getData());
      }
      unknownFields.writeTo(output);
    }

    @java.lang.Override
    public int getSerializedSize() {
      int size = memoizedSize;
      if (size != -1) return size;

      size = 0;
      if (data_ != null) {
        size += com.google.protobuf.CodedOutputStream
          .computeMessageSize(1, getData());
      }
      size += unknownFields.getSerializedSize();
      memoizedSize = size;
      return size;
    }

    @java.lang.Override
    public boolean equals(final java.lang.Object obj) {
      if (obj == this) {
       return true;
      }
      if (!(obj instanceof UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest)) {
        return super.equals(obj);
      }
      UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest other = (UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest) obj;

      if (hasData() != other.hasData()) return false;
      if (hasData()) {
        if (!getData()
            .equals(other.getData())) return false;
      }
      if (!unknownFields.equals(other.unknownFields)) return false;
      return true;
    }

    @java.lang.Override
    public int hashCode() {
      if (memoizedHashCode != 0) {
        return memoizedHashCode;
      }
      int hash = 41;
      hash = (19 * hash) + getDescriptor().hashCode();
      if (hasData()) {
        hash = (37 * hash) + DATA_FIELD_NUMBER;
        hash = (53 * hash) + getData().hashCode();
      }
      hash = (29 * hash) + unknownFields.hashCode();
      memoizedHashCode = hash;
      return hash;
    }

    public static UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest parseFrom(
        java.nio.ByteBuffer data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest parseFrom(
        java.nio.ByteBuffer data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest parseFrom(
        com.google.protobuf.ByteString data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest parseFrom(
        com.google.protobuf.ByteString data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest parseFrom(byte[] data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest parseFrom(
        byte[] data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest parseFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest parseFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input, extensionRegistry);
    }
    public static UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest parseDelimitedFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input);
    }
    public static UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest parseDelimitedFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
    }
    public static UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest parseFrom(
        com.google.protobuf.CodedInputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest parseFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input, extensionRegistry);
    }

    @java.lang.Override
    public Builder newBuilderForType() { return newBuilder(); }
    public static Builder newBuilder() {
      return DEFAULT_INSTANCE.toBuilder();
    }
    public static Builder newBuilder(UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest prototype) {
      return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
    }
    @java.lang.Override
    public Builder toBuilder() {
      return this == DEFAULT_INSTANCE
          ? new Builder() : new Builder().mergeFrom(this);
    }

    @java.lang.Override
    protected Builder newBuilderForType(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      Builder builder = new Builder(parent);
      return builder;
    }
    /**
     * Protobuf type {@code UniversaltestService.CreateUniversaltestServiceRequest}
     */
    public static final class Builder extends
        com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
        // @@protoc_insertion_point(builder_implements:UniversaltestService.CreateUniversaltestServiceRequest)
        UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequestOrBuilder {
      public static final com.google.protobuf.Descriptors.Descriptor
          getDescriptor() {
        return UniversaltestService.Universaltestservice.internal_static_UniversaltestService_CreateUniversaltestServiceRequest_descriptor;
      }

      @java.lang.Override
      protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
          internalGetFieldAccessorTable() {
        return UniversaltestService.Universaltestservice.internal_static_UniversaltestService_CreateUniversaltestServiceRequest_fieldAccessorTable
            .ensureFieldAccessorsInitialized(
                UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest.class, UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest.Builder.class);
      }

      // Construct using UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest.newBuilder()
      private Builder() {
        maybeForceBuilderInitialization();
      }

      private Builder(
          com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
        super(parent);
        maybeForceBuilderInitialization();
      }
      private void maybeForceBuilderInitialization() {
        if (com.google.protobuf.GeneratedMessageV3
                .alwaysUseFieldBuilders) {
        }
      }
      @java.lang.Override
      public Builder clear() {
        super.clear();
        if (dataBuilder_ == null) {
          data_ = null;
        } else {
          data_ = null;
          dataBuilder_ = null;
        }
        return this;
      }

      @java.lang.Override
      public com.google.protobuf.Descriptors.Descriptor
          getDescriptorForType() {
        return UniversaltestService.Universaltestservice.internal_static_UniversaltestService_CreateUniversaltestServiceRequest_descriptor;
      }

      @java.lang.Override
      public UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest getDefaultInstanceForType() {
        return UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest.getDefaultInstance();
      }

      @java.lang.Override
      public UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest build() {
        UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest result = buildPartial();
        if (!result.isInitialized()) {
          throw newUninitializedMessageException(result);
        }
        return result;
      }

      @java.lang.Override
      public UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest buildPartial() {
        UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest result = new UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest(this);
        if (dataBuilder_ == null) {
          result.data_ = data_;
        } else {
          result.data_ = dataBuilder_.build();
        }
        onBuilt();
        return result;
      }

      @java.lang.Override
      public Builder clone() {
        return super.clone();
      }
      @java.lang.Override
      public Builder setField(
          com.google.protobuf.Descriptors.FieldDescriptor field,
          java.lang.Object value) {
        return super.setField(field, value);
      }
      @java.lang.Override
      public Builder clearField(
          com.google.protobuf.Descriptors.FieldDescriptor field) {
        return super.clearField(field);
      }
      @java.lang.Override
      public Builder clearOneof(
          com.google.protobuf.Descriptors.OneofDescriptor oneof) {
        return super.clearOneof(oneof);
      }
      @java.lang.Override
      public Builder setRepeatedField(
          com.google.protobuf.Descriptors.FieldDescriptor field,
          int index, java.lang.Object value) {
        return super.setRepeatedField(field, index, value);
      }
      @java.lang.Override
      public Builder addRepeatedField(
          com.google.protobuf.Descriptors.FieldDescriptor field,
          java.lang.Object value) {
        return super.addRepeatedField(field, value);
      }
      @java.lang.Override
      public Builder mergeFrom(com.google.protobuf.Message other) {
        if (other instanceof UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest) {
          return mergeFrom((UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest)other);
        } else {
          super.mergeFrom(other);
          return this;
        }
      }

      public Builder mergeFrom(UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest other) {
        if (other == UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest.getDefaultInstance()) return this;
        if (other.hasData()) {
          mergeData(other.getData());
        }
        this.mergeUnknownFields(other.unknownFields);
        onChanged();
        return this;
      }

      @java.lang.Override
      public final boolean isInitialized() {
        return true;
      }

      @java.lang.Override
      public Builder mergeFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws java.io.IOException {
        UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest parsedMessage = null;
        try {
          parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
        } catch (com.google.protobuf.InvalidProtocolBufferException e) {
          parsedMessage = (UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest) e.getUnfinishedMessage();
          throw e.unwrapIOException();
        } finally {
          if (parsedMessage != null) {
            mergeFrom(parsedMessage);
          }
        }
        return this;
      }

      private universaltest.UniversaltestOuterClass.Universaltest data_;
      private com.google.protobuf.SingleFieldBuilderV3<
          universaltest.UniversaltestOuterClass.Universaltest, universaltest.UniversaltestOuterClass.Universaltest.Builder, universaltest.UniversaltestOuterClass.UniversaltestOrBuilder> dataBuilder_;
      /**
       * <code>.universaltest.Universaltest data = 1;</code>
       */
      public boolean hasData() {
        return dataBuilder_ != null || data_ != null;
      }
      /**
       * <code>.universaltest.Universaltest data = 1;</code>
       */
      public universaltest.UniversaltestOuterClass.Universaltest getData() {
        if (dataBuilder_ == null) {
          return data_ == null ? universaltest.UniversaltestOuterClass.Universaltest.getDefaultInstance() : data_;
        } else {
          return dataBuilder_.getMessage();
        }
      }
      /**
       * <code>.universaltest.Universaltest data = 1;</code>
       */
      public Builder setData(universaltest.UniversaltestOuterClass.Universaltest value) {
        if (dataBuilder_ == null) {
          if (value == null) {
            throw new NullPointerException();
          }
          data_ = value;
          onChanged();
        } else {
          dataBuilder_.setMessage(value);
        }

        return this;
      }
      /**
       * <code>.universaltest.Universaltest data = 1;</code>
       */
      public Builder setData(
          universaltest.UniversaltestOuterClass.Universaltest.Builder builderForValue) {
        if (dataBuilder_ == null) {
          data_ = builderForValue.build();
          onChanged();
        } else {
          dataBuilder_.setMessage(builderForValue.build());
        }

        return this;
      }
      /**
       * <code>.universaltest.Universaltest data = 1;</code>
       */
      public Builder mergeData(universaltest.UniversaltestOuterClass.Universaltest value) {
        if (dataBuilder_ == null) {
          if (data_ != null) {
            data_ =
              universaltest.UniversaltestOuterClass.Universaltest.newBuilder(data_).mergeFrom(value).buildPartial();
          } else {
            data_ = value;
          }
          onChanged();
        } else {
          dataBuilder_.mergeFrom(value);
        }

        return this;
      }
      /**
       * <code>.universaltest.Universaltest data = 1;</code>
       */
      public Builder clearData() {
        if (dataBuilder_ == null) {
          data_ = null;
          onChanged();
        } else {
          data_ = null;
          dataBuilder_ = null;
        }

        return this;
      }
      /**
       * <code>.universaltest.Universaltest data = 1;</code>
       */
      public universaltest.UniversaltestOuterClass.Universaltest.Builder getDataBuilder() {
        
        onChanged();
        return getDataFieldBuilder().getBuilder();
      }
      /**
       * <code>.universaltest.Universaltest data = 1;</code>
       */
      public universaltest.UniversaltestOuterClass.UniversaltestOrBuilder getDataOrBuilder() {
        if (dataBuilder_ != null) {
          return dataBuilder_.getMessageOrBuilder();
        } else {
          return data_ == null ?
              universaltest.UniversaltestOuterClass.Universaltest.getDefaultInstance() : data_;
        }
      }
      /**
       * <code>.universaltest.Universaltest data = 1;</code>
       */
      private com.google.protobuf.SingleFieldBuilderV3<
          universaltest.UniversaltestOuterClass.Universaltest, universaltest.UniversaltestOuterClass.Universaltest.Builder, universaltest.UniversaltestOuterClass.UniversaltestOrBuilder> 
          getDataFieldBuilder() {
        if (dataBuilder_ == null) {
          dataBuilder_ = new com.google.protobuf.SingleFieldBuilderV3<
              universaltest.UniversaltestOuterClass.Universaltest, universaltest.UniversaltestOuterClass.Universaltest.Builder, universaltest.UniversaltestOuterClass.UniversaltestOrBuilder>(
                  getData(),
                  getParentForChildren(),
                  isClean());
          data_ = null;
        }
        return dataBuilder_;
      }
      @java.lang.Override
      public final Builder setUnknownFields(
          final com.google.protobuf.UnknownFieldSet unknownFields) {
        return super.setUnknownFields(unknownFields);
      }

      @java.lang.Override
      public final Builder mergeUnknownFields(
          final com.google.protobuf.UnknownFieldSet unknownFields) {
        return super.mergeUnknownFields(unknownFields);
      }


      // @@protoc_insertion_point(builder_scope:UniversaltestService.CreateUniversaltestServiceRequest)
    }

    // @@protoc_insertion_point(class_scope:UniversaltestService.CreateUniversaltestServiceRequest)
    private static final UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest DEFAULT_INSTANCE;
    static {
      DEFAULT_INSTANCE = new UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest();
    }

    public static UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest getDefaultInstance() {
      return DEFAULT_INSTANCE;
    }

    private static final com.google.protobuf.Parser<CreateUniversaltestServiceRequest>
        PARSER = new com.google.protobuf.AbstractParser<CreateUniversaltestServiceRequest>() {
      @java.lang.Override
      public CreateUniversaltestServiceRequest parsePartialFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws com.google.protobuf.InvalidProtocolBufferException {
        return new CreateUniversaltestServiceRequest(input, extensionRegistry);
      }
    };

    public static com.google.protobuf.Parser<CreateUniversaltestServiceRequest> parser() {
      return PARSER;
    }

    @java.lang.Override
    public com.google.protobuf.Parser<CreateUniversaltestServiceRequest> getParserForType() {
      return PARSER;
    }

    @java.lang.Override
    public UniversaltestService.Universaltestservice.CreateUniversaltestServiceRequest getDefaultInstanceForType() {
      return DEFAULT_INSTANCE;
    }

  }

  private static final com.google.protobuf.Descriptors.Descriptor
    internal_static_UniversaltestService_CreateUniversaltestServiceRequest_descriptor;
  private static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_UniversaltestService_CreateUniversaltestServiceRequest_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n/UniversaltestService/universaltestserv" +
      "ice.proto\022\024UniversaltestService\032\034google/" +
      "api/annotations.proto\032!universaltest/uni" +
      "versaltest.proto\032\033google/protobuf/empty." +
      "proto\"O\n!CreateUniversaltestServiceReque" +
      "st\022*\n\004data\030\001 \001(\0132\034.universaltest.Univers" +
      "altest2\263\001\n\024UniversaltestService\022\232\001\n\023Crea" +
      "teuniversaltest\0227.UniversaltestService.C" +
      "reateUniversaltestServiceRequest\032\".unive" +
      "rsaltest.UniversaltestEntity\"&\202\323\344\223\002 \"\030/m" +
      "ockdata/universaltests:\004dataB\027Z\025/Univers" +
      "altestServiceb\006proto3"
    };
    com.google.protobuf.Descriptors.FileDescriptor.InternalDescriptorAssigner assigner =
        new com.google.protobuf.Descriptors.FileDescriptor.    InternalDescriptorAssigner() {
          public com.google.protobuf.ExtensionRegistry assignDescriptors(
              com.google.protobuf.Descriptors.FileDescriptor root) {
            descriptor = root;
            return null;
          }
        };
    com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          com.google.api.AnnotationsProto.getDescriptor(),
          universaltest.UniversaltestOuterClass.getDescriptor(),
          google.protobuf.EmptyOuterClass.getDescriptor(),
        }, assigner);
    internal_static_UniversaltestService_CreateUniversaltestServiceRequest_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_UniversaltestService_CreateUniversaltestServiceRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_UniversaltestService_CreateUniversaltestServiceRequest_descriptor,
        new java.lang.String[] { "Data", });
    com.google.protobuf.ExtensionRegistry registry =
        com.google.protobuf.ExtensionRegistry.newInstance();
    registry.add(com.google.api.AnnotationsProto.http);
    com.google.protobuf.Descriptors.FileDescriptor
        .internalUpdateFileDescriptor(descriptor, registry);
    com.google.api.AnnotationsProto.getDescriptor();
    universaltest.UniversaltestOuterClass.getDescriptor();
    google.protobuf.EmptyOuterClass.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
