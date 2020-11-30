<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: graphik.proto

namespace Api;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>api.Authorizers</code>
 */
class Authorizers extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated .api.Authorizer authorizers = 1;</code>
     */
    private $authorizers;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Api\Authorizer[]|\Google\Protobuf\Internal\RepeatedField $authorizers
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Graphik::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated .api.Authorizer authorizers = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getAuthorizers()
    {
        return $this->authorizers;
    }

    /**
     * Generated from protobuf field <code>repeated .api.Authorizer authorizers = 1;</code>
     * @param \Api\Authorizer[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setAuthorizers($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Api\Authorizer::class);
        $this->authorizers = $arr;

        return $this;
    }

}

