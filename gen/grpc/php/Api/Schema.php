<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: graphik.proto

namespace Api;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Schema returns registered connection & doc types
 *
 * Generated from protobuf message <code>api.Schema</code>
 */
class Schema extends \Google\Protobuf\Internal\Message
{
    /**
     * connection_types are the types of connections in the graph
     *
     * Generated from protobuf field <code>repeated string connection_types = 1;</code>
     */
    private $connection_types;
    /**
     * doc_types are the types of docs in the graph
     *
     * Generated from protobuf field <code>repeated string doc_types = 2;</code>
     */
    private $doc_types;
    /**
     * Generated from protobuf field <code>.api.Authorizers authorizers = 3;</code>
     */
    private $authorizers = null;
    /**
     * Generated from protobuf field <code>.api.Constraints constraints = 4;</code>
     */
    private $constraints = null;
    /**
     * Generated from protobuf field <code>.api.Indexes indexes = 5;</code>
     */
    private $indexes = null;
    /**
     * Generated from protobuf field <code>.api.Triggers triggers = 6;</code>
     */
    private $triggers = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $connection_types
     *           connection_types are the types of connections in the graph
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $doc_types
     *           doc_types are the types of docs in the graph
     *     @type \Api\Authorizers $authorizers
     *     @type \Api\Constraints $constraints
     *     @type \Api\Indexes $indexes
     *     @type \Api\Triggers $triggers
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Graphik::initOnce();
        parent::__construct($data);
    }

    /**
     * connection_types are the types of connections in the graph
     *
     * Generated from protobuf field <code>repeated string connection_types = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getConnectionTypes()
    {
        return $this->connection_types;
    }

    /**
     * connection_types are the types of connections in the graph
     *
     * Generated from protobuf field <code>repeated string connection_types = 1;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setConnectionTypes($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->connection_types = $arr;

        return $this;
    }

    /**
     * doc_types are the types of docs in the graph
     *
     * Generated from protobuf field <code>repeated string doc_types = 2;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getDocTypes()
    {
        return $this->doc_types;
    }

    /**
     * doc_types are the types of docs in the graph
     *
     * Generated from protobuf field <code>repeated string doc_types = 2;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setDocTypes($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->doc_types = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.api.Authorizers authorizers = 3;</code>
     * @return \Api\Authorizers
     */
    public function getAuthorizers()
    {
        return $this->authorizers;
    }

    /**
     * Generated from protobuf field <code>.api.Authorizers authorizers = 3;</code>
     * @param \Api\Authorizers $var
     * @return $this
     */
    public function setAuthorizers($var)
    {
        GPBUtil::checkMessage($var, \Api\Authorizers::class);
        $this->authorizers = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.api.Constraints constraints = 4;</code>
     * @return \Api\Constraints
     */
    public function getConstraints()
    {
        return $this->constraints;
    }

    /**
     * Generated from protobuf field <code>.api.Constraints constraints = 4;</code>
     * @param \Api\Constraints $var
     * @return $this
     */
    public function setConstraints($var)
    {
        GPBUtil::checkMessage($var, \Api\Constraints::class);
        $this->constraints = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.api.Indexes indexes = 5;</code>
     * @return \Api\Indexes
     */
    public function getIndexes()
    {
        return $this->indexes;
    }

    /**
     * Generated from protobuf field <code>.api.Indexes indexes = 5;</code>
     * @param \Api\Indexes $var
     * @return $this
     */
    public function setIndexes($var)
    {
        GPBUtil::checkMessage($var, \Api\Indexes::class);
        $this->indexes = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.api.Triggers triggers = 6;</code>
     * @return \Api\Triggers
     */
    public function getTriggers()
    {
        return $this->triggers;
    }

    /**
     * Generated from protobuf field <code>.api.Triggers triggers = 6;</code>
     * @param \Api\Triggers $var
     * @return $this
     */
    public function setTriggers($var)
    {
        GPBUtil::checkMessage($var, \Api\Triggers::class);
        $this->triggers = $var;

        return $this;
    }

}

