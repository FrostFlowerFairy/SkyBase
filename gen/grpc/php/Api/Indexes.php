<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: graphik.proto

namespace Api;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>api.Indexes</code>
 */
class Indexes extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated .api.Index indexes = 1;</code>
     */
    private $indexes;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Api\Index[]|\Google\Protobuf\Internal\RepeatedField $indexes
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Graphik::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated .api.Index indexes = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getIndexes()
    {
        return $this->indexes;
    }

    /**
     * Generated from protobuf field <code>repeated .api.Index indexes = 1;</code>
     * @param \Api\Index[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setIndexes($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Api\Index::class);
        $this->indexes = $arr;

        return $this;
    }

}

